package main

import (
	"context"
	"io"
	"time"

	"github.com/Jorrit05/DYNAMOS/pkg/lib"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
)

func startConsumingWithRetry(c pb.SideCarClient, name string, maxRetries int, waitTime time.Duration) {
	for i := 0; i < maxRetries; i++ {
		err := startConsuming(c, name)
		if err == nil {
			return
		}

		logger.Sugar().Errorf("Failed to start consuming (attempt %d/%d): %v", i+1, maxRetries, err)

		// Wait for some time before retrying
		time.Sleep(waitTime)
	}
}

func startConsuming(c pb.SideCarClient, from string) error {
	logger.Debug("Start startConsuming function")
	ctx := context.Background()
	stream, err := c.Consume(ctx, &pb.ConsumeRequest{QueueName: from, AutoAck: true})
	if err != nil {
		logger.Sugar().Fatalf("Error on consume: %v", err)
	}

	for {
		grpcMsg, err := stream.Recv()
		if err == io.EOF {
			// The stream has ended.
			logger.Sugar().Warnw("Stream has ended", "error:", err)
			break
		}

		if err != nil {
			logger.Sugar().Fatalf("Failed to receive: %v", err)
		}

		ctx, span, err := lib.StartRemoteParentSpan(ctx, serviceName+"/consume/"+grpcMsg.Type, grpcMsg.Trace)
		if err != nil {
			logger.Sugar().Errorf("Error starting span: %v", err)
			return err
		}

		switch grpcMsg.Type {
		case "requestApproval":
			logger.Debug("sidecar/Consume: Received a requestApproval")
			var requestApproval pb.RequestApproval
			if err := grpcMsg.Body.UnmarshalTo(&requestApproval); err != nil {
				logger.Sugar().Fatalf("Failed to unmarshal message: %v", err)
			}

			logger.Sugar().Infof("User name: %s", requestApproval.User.UserName)
			checkRequestApproval(ctx, &requestApproval)

		default:

			logger.Sugar().Fatalf("Unknown message type: %s", grpcMsg.Type)
		}
		logger.Sugar().Warnf("ENDING, SPAN PE:  %v", time.Now())

		span.End()
	}
	return err
}
