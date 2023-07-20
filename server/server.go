package main

import (
	"context"
	"fmt"
	"log"
	"net"
	
	pb "example.com/starter/proto/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (*server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Printf("Received: %v", req.GetName())
	return &pb.HelloReply{Message: "Hello " + req.GetName()}, nil
}

func (*server) mustEmbedUnimplementedGreeterServer() {}

func main() {
	fmt.Println("Starting gRPC server...")
	
	// Listen to port 50051
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	
	// Create a server instance
	grpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(grpcServer, &server{})  // what is &server{} means
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	
}