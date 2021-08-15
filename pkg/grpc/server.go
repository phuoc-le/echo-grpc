package grpc

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-echo/pkg/proto"
	"log"
)

type server struct {
	pb.UnimplementedGreeterServiceServer
	pb.UnimplementedTestServiceServer
}

// Test
func (s *server) TestEcho(ctx context.Context, in *pb.TestRequest) (*pb.TestResponse, error) {
	log.Printf("Test: %v", in.GetMessage())
	return &pb.TestResponse{Message: "Hello " + in.GetMessage()}, nil
}

// SayHello
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func Init(s *grpc.Server) {
	pb.RegisterGreeterServiceServer(s, &server{})
	pb.RegisterTestServiceServer(s, &server{})
}
