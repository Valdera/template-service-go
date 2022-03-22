package handler

import (
	"context"
	"test/core/module"
	pb "test/pb"

	"github.com/golang/protobuf/ptypes/empty"
)

type testHandler struct {
	testUc module.TestUseCase
	pb.UnimplementedTestServer
}

func NewTestHandler(testUc module.TestUseCase) *testHandler {
	return &testHandler{
		testUc: testUc,
	}
}

func (s *testHandler) GetTest(ctx context.Context, req *empty.Empty) (*pb.TestResponse, error) {
	return &pb.TestResponse{
		Message: "hi",
	}, nil
}
