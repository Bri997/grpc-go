package main

import (
	"context"
	"fmt"
	"log"

	"github.com/bri997/grpc-go-course/greet/hands-on/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hi cal client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)

	}
	defer cc.Close()

	c := calculatorpb.NewCalculateServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculateServiceClient) {
	fmt.Println("starting to do a Unary RPC...")
	req := &calculatorpb.SumRequest{
		Num1: 5,
		Num2: 10,
	}
	res, err := c.Sum(context.Background(), req)
	if err != nil {
		println("req =", req)
		log.Fatalf("error while calling Sum RPC: %v", err)
	}
	log.Printf("Response from Sum: %v", res.Result)
}
