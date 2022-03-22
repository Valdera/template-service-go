package main

import (
	"test/config"
	"test/core/module"
	"test/core/repository"

	"google.golang.org/grpc"

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
}
