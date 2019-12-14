package main

import (
	"context"
	"fmt"
	"io"
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

	//doUnary(c)
	//doStreming(c)
	doClientStreaming(c)
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

func doStreming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting the server Prime stream...")
	req := &calculatorpb.CalcPrimeRequest{

		Number: 156546512,
	}

	stream, err := c.CalcPrime(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CalcPrime", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Problem for loop", err)
		}
		fmt.Println(res.GetPrimeFactor())

	}

}

func doClientStreaming(c calculatorpb.CalculateServiceClient) {
	fmt.Println("Starting the Client Avg server Prime stream...")
	stream, err := c.CalcAverage(context.Background())
	if err != nil {
		log.Fatalf("Err while opening stream ", err)
	}
	numbers := []int32{3, 5, 9, 54, 23}

	for _, number := range numbers {
		stream.Send(&calculatorpb.CalcAvgRequest{
			Number: number,
		})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response:..", err)
	}
	fmt.Printf("The average is: %v\n", res.GetAveResult())
}
