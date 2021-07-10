package main

import (
	"context"
	"fmt"
	"grpc-course/greet/greetpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("greeet request was invoked with%v", req)
	firstname := req.GetGreeting().GetFirstName()
	result := "Hello" + firstname
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil

}

func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal("Failed to Listen %v", err)
	}
	s := grpc.NewServer()
	s1 := &server{}
	greetpb.RegisterGreetServiceServer(s, s1)

	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve %v", err)
	}
}
