package main

import (
	"fmt"
	"grpc-course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to Listen %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, greetpb.UnimplementedGreetServiceServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %v", err)
	}
}
