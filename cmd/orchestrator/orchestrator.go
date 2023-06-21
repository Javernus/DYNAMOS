package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/Jorrit05/DYNAMOS/pkg/lib"
	pb "github.com/Jorrit05/DYNAMOS/pkg/proto"
	"github.com/gorilla/handlers"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
)

var (
	logger                      = lib.InitLogger()
	etcdClient *clientv3.Client = lib.GetEtcdClient(etcdEndpoints)
	c          pb.SideCarClient
	conn       *grpc.ClientConn
)

func main() {
	defer logger.Sync() // flushes buffer, if any
	defer etcdClient.Close()

	c, conn = lib.InitializeRabbit(grpcAddr, &pb.ServiceRequest{ServiceName: fmt.Sprintf("%s-in", serviceName), RoutingKey: fmt.Sprintf("%s-in", serviceName), QueueAutoDelete: false})
	defer conn.Close()

	// Define a WaitGroup
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		startConsumingWithRetry(c, fmt.Sprintf("%s-in", serviceName), 5, 5*time.Second)

		wg.Done() // Decrement the WaitGroup counter when the goroutine finishes
	}()

	go registerPolicyEnforcerConfiguration()
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	mux := http.NewServeMux()

	apiMux := http.NewServeMux()
	apiMux.HandleFunc("/archetypes", archetypesHandler(etcdClient, "/archetypes"))
	apiMux.HandleFunc("/archetypes/", archetypesHandler(etcdClient, "/archetypes"))

	apiMux.HandleFunc("/requesttypes", requestTypesHandler(etcdClient, "/requestTypes"))
	apiMux.HandleFunc("/requestTypes", requestTypesHandler(etcdClient, "/requestTypes"))

	apiMux.HandleFunc("/requesttypes/", requestTypesHandler(etcdClient, "/requestTypes"))
	apiMux.HandleFunc("/requestTypes/", requestTypesHandler(etcdClient, "/requestTypes"))

	apiMux.HandleFunc("/microservices", microserviceMetadataHandler(etcdClient, "/microservices"))
	apiMux.HandleFunc("/microservices/", microserviceMetadataHandler(etcdClient, "/microservices"))

	apiMux.HandleFunc("/policyEnforcer", agreementsHandler(etcdClient, "/policyEnforcer"))
	apiMux.HandleFunc("/policyEnforcer/", agreementsHandler(etcdClient, "/policyEnforcer"))

	apiMux.HandleFunc("/requestapproval", requestApprovalHandler())
	logger.Info(apiVersion) // prints /api/v1

	mux.Handle(apiVersion+"/", http.StripPrefix(apiVersion, apiMux))

	logger.Sugar().Infow("Starting http server on: ", "port", port)
	go func() {
		if err := http.ListenAndServe(port, logMiddleware(handlers.CORS(originsOk, headersOk, methodsOk)(mux))); err != nil {
			logger.Sugar().Fatalw("Error starting HTTP server: %s", err)
		}
	}()

	// go func() {
	// 	if err := http.ListenAndServe(port, handlers.CORS(originsOk, headersOk, methodsOk)(mux)); err != nil {
	// 		logger.Sugar().Fatalw("Error starting HTTP server: %s", err)
	// 	}
	// }()

	wg.Wait()

}

func logMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		logger.Sugar().Infow("Received request", "method", r.Method, "path", r.URL.Path)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
