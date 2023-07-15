package main

import (
	"context"
	"errors"
	"time"

	"github.com/Jorrit05/DYNAMOS/pkg/etcd"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	bo "github.com/cenkalti/backoff/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	channel    *amqp.Channel
	conn       *amqp.Connection
	messages   <-chan amqp.Delivery
	routingKey string
)

func send(ctx context.Context, message amqp.Publishing, target string, opts ...etcd.Option) (*emptypb.Empty, error) {
	// Start with default options
	ctx, span := trace.StartSpan(ctx, "send/"+target)

	retryOpts := etcd.DefaultRetryOptions

	// Apply any specified options
	for _, opt := range opts {
		opt(&retryOpts)
	}

	// Create a returns channel
	returns := make(chan amqp.Return)
	channel.NotifyReturn(returns)

	sc := trace.FromContext(ctx).SpanContext()
	binarySc := propagation.Binary(sc)

	if message.Headers == nil {
		message.Headers = amqp.Table{}
	}

	message.Headers["trace"] = binarySc

	operation := func() error {
		// Log before sending message
		logger.Sugar().Infow("Sending message: ", "My routingKey", routingKey, "exchangeName", exchangeName, "target", target)

		errChan := make(chan error)

		go func() {
			// Create a context with a timeout
			timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()

			// _, span := trace.StartSpan(timeoutCtx, "publish")
			err := channel.PublishWithContext(timeoutCtx, exchangeName, target, true, false, message)
			errChan <- err
		}()

		select {
		case err := <-errChan:
			if err != nil {
				logger.Sugar().Debugf("In error chan: %v", err)
				return err
			}
		case r := <-returns:
			if r.ReplyText == "NO_ROUTE" {
				logger.Sugar().Infof("No route to job yet: %v", target)
				// This will trigger a retry
				return errors.New("no route to target")
			} else {
				logger.Sugar().Errorf("Unknown reason message returned: %v", r)
				return bo.Permanent(errors.New("unknown error"))
			}
		case <-time.After(3 * time.Second): // Timeout if no message is received in 3 seconds
			logger.Sugar().Debugf("No message received in 3 seconds")
		}

		return nil
	}

	// Create a new exponential backoff
	backoff := bo.NewExponentialBackOff()
	backoff.InitialInterval = retryOpts.InitialInterval
	backoff.MaxInterval = retryOpts.MaxInterval
	backoff.MaxElapsedTime = retryOpts.MaxElapsedTime
	span.End()

	err := bo.Retry(operation, backoff)
	if err != nil {
		logger.Sugar().Errorf("Publish failed after %v seconds, err: %s", backoff.MaxElapsedTime, err)
		return &emptypb.Empty{}, err
	}
	return &emptypb.Empty{}, nil
}

func (s *server) SendRequestApproval(ctx context.Context, in *pb.RequestApproval) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal requestApproval failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "requestApproval",
	}

	return send(ctx, message, "policyEnforcer-in")
}

func (s *server) SendValidationResponse(ctx context.Context, in *pb.ValidationResponse) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal ValidationResponse failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "validationResponse",
	}

	return send(ctx, message, "orchestrator-in")
}

func (s *server) SendCompositionRequest(ctx context.Context, in *pb.CompositionRequest) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal CompositionRequest failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		Body: data,
		Type: "compositionRequest",
	}

	return send(ctx, message, in.DestinationQueue)
}

func (s *server) SendSqlDataRequest(ctx context.Context, in *pb.SqlDataRequest) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal requestApproval failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		CorrelationId: in.RequestMetada.CorrelationId,
		Body:          data,
		Type:          "sqlDataRequest",
	}
	logger.Sugar().Debugf("SendSqlDataRequest destination queue: %v", in.RequestMetada.DestinationQueue)
	return send(ctx, message, in.RequestMetada.DestinationQueue, etcd.WithMaxElapsedTime(10*time.Second))
}

func (s *server) SendMicroserviceComm(ctx context.Context, in *pb.MicroserviceCommunication) (*emptypb.Empty, error) {
	data, err := proto.Marshal(in)
	if err != nil {
		logger.Sugar().Errorf("Marshal requestApproval failed: %s", err)

		return nil, status.Error(codes.Internal, err.Error())
	}

	// Do other stuff
	message := amqp.Publishing{
		CorrelationId: in.CorrelationId,
		Body:          data,
		Type:          in.Type,
	}

	return send(ctx, message, in.DestinationQueue, etcd.WithMaxElapsedTime(10*time.Second))
}

// func (s *server) SendSqlDataRequestResponse(ctx context.Context, in *pb.SqlDataRequestResponse) (*emptypb.Empty, error) {
// 	data, err := proto.Marshal(in)
// 	if err != nil {
// 		logger.Sugar().Errorf("Marshal requestApproval failed: %s", err)

// 		return nil, status.Error(codes.Internal, err.Error())
// 	}

// 	// Do other stuff
// 	message := amqp.Publishing{
// 		CorrelationId: in.CorrelationId,
// 		Body:          data,
// 		Type:          "sqlDataRequestResponse",
// 	}

// 	return send(ctx, message, in.DestinationQueue, etcd.WithMaxElapsedTime(10*time.Second))
// }
