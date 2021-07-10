package main

import (
	"fmt"
	"grpc-course/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hello im client :")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)

	}
	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("created client :%f", c)

}
