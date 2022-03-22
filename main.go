package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"test/config"
	"test/core/module"
	"test/core/repository"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	handler "test/handler"
	pb "test/pb"
)

func main() {
	cfg := config.Get()

	s := grpc.NewServer()
	testRepo := repository.NewTestRepository(cfg)
	testUc := module.NewTestUsecase(cfg, testRepo)
	testHandler := handler.NewTestHandler(testUc)

	pb.RegisterTestServer(s, testHandler)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go serveGRPC(s, cfg)
	go serveREST(cfg)

	<-done

	s.GracefulStop()

	log.Print("Server Exited Properly")
}

func serveGRPC(s *grpc.Server, cfg config.Config) {
	lis, err := net.Listen("tcp", cfg.GRPCPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	reflection.Register(s)
	log.Println("gRPC services listening on port ", cfg.GRPCPort)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func serveREST(cfg config.Config) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := pb.RegisterTestHandlerFromEndpoint(ctx, mux, cfg.HOST+cfg.GRPCPort, opts)
	if err != nil {
		log.Fatalf("Fail to register gRPC service endpoint: %v", err)
		return
	}

	log.Println("REST services listening on port ", cfg.RESTPort)
	if err := http.ListenAndServe(cfg.RESTPort, mux); err != nil {
		log.Fatalf("Could not setup HTTP endpoint: %v", err)
	}

}
