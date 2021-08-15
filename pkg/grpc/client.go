package grpc

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-echo/pkg/proto"
	"github.com/labstack/gommon/log"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)
func client() *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func GreeterService() {

	defer client().Close()
	greeterCnn := pb.NewGreeterServiceClient(client())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	greeterRs, err := greeterCnn.SayHello(ctx, &pb.HelloRequest{Name: "Phuoc"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", greeterRs.GetMessage())
}

func TestService(name string) {
	defer client().Close()
	testCnn := pb.NewTestServiceClient(client())
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	testRs, err := testCnn.TestEcho(ctx, &pb.TestRequest{Message: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", testRs.GetMessage())
}