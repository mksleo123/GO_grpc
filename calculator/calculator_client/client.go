package main

import (
	"context"
	"fmt"
	"grpc-course/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator  client :")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal("could not connect: %v", err)

	}
	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)
	//fmt.Printf("created client :%f", c)
	doUnary(c)

}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("starting to do unary rpc")
	req := &calculatorpb.SumRequest{
		FirstNumber:   2,
		SeconfdNumber: 5,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatal("error while calling sum rpc %v", err)
	}
	log.Printf("Rresponse froom greet  %v", res.SumResult)
}
