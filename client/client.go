package main

import (
	"context"
	"log"
	"time"
	
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "example.com/starter/proto/helloworld"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect")
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Cowbon"})
	if err != nil {
		log.Fatalf("could not greet")
	}
	log.Printf("Greeting: %s", r.GetMessage())
}