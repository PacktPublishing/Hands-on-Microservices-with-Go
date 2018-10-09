package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/grpc/pb"

	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/repositories"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/service"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/endpoints"
	grpcEncoding "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/grpc/encoding"
	httpEncoding "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/http/encoding"
	"github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/middleware"
	"google.golang.org/grpc"

	gokitLog "github.com/go-kit/kit/log"

	mygrpctransport "github.com/PacktPublishing/Hands-on-Microservices-with-Go/section-8/video-2/agents-service/src/api/transports/grpc/transport"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/gorilla/mux"
)

func main() {

	logger := gokitLog.NewLogfmtLogger(os.Stderr)

	repo := repositories.NewMariaDBAgentsRepository()
	defer repo.Close()

	svc := service.AgentsServiceImpl{}
	svc.Repo = repo

	loggingMiddleware := middleware.LoggingMiddleware(gokitLog.With(logger, "method", "insertAgentPlayer"))

	insertAgentPlayerEndpoint := endpoints.MakeInsertAgentPlayerEndpoint(svc)
	insertAgentPlayerEndpoint = loggingMiddleware(insertAgentPlayerEndpoint)

	getAgentByIDEndpoint := endpoints.MakeGetAgentByIDEndpoint(svc)
	getAgentByIDEndpoint = middleware.LoggingMiddleware(gokitLog.With(logger, "method", "getAgent"))(getAgentByIDEndpoint)

	//GRPC

	gsrv := &mygrpctransport.GrpcServer{
		InsertAgentPlayerHandler: grpctransport.NewServer(
			insertAgentPlayerEndpoint,
			grpcEncoding.DecodeInsertAgentPlayerRequest,
			grpcEncoding.EncodeInsertAgentPlayerResponse,
		),
		GetAgentByIDHandler: grpctransport.NewServer(
			getAgentByIDEndpoint,
			grpcEncoding.DecodeGetAgentByIDRequest,
			grpcEncoding.EncodeGetAgentByIDResponse,
		),
	}

	// The gRPC listener mounts the Go kit gRPC server we created.
	grpcListener, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Log("transport", "gRPC", "during", "Listen", "err", err)
		os.Exit(1)
	}
	defer grpcListener.Close()

	logger.Log("transport", "gRPC", "addr", "50051")
	s := grpc.NewServer()
	pb.RegisterAgentsServer(s, gsrv)
	go func() {
		log.Fatal(s.Serve(grpcListener))
	}()

	//HTTP
	r := mux.NewRouter()

	r.Methods("POST").Path("/agent-player").Handler(httptransport.NewServer(
		insertAgentPlayerEndpoint,
		httpEncoding.DecodeInsertAgentPlayerRequest,
		httpEncoding.EncodeInsertAgentPlayerResponse,
	))

	r.Methods("GET").Path("/agent/{id}").Handler(httptransport.NewServer(
		getAgentByIDEndpoint,
		httpEncoding.DecodeGetAgentByIDRequest,
		httpEncoding.EncodeGetAgentByIDResponse,
	))

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Starting server on port: 8080")
	log.Fatal(srv.ListenAndServe())

}
