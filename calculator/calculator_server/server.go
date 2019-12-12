package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/bri997/grpc-go-course/greet/hands-on/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {

	fmt.Printf("Cal function was invoked with %v\n", req)
	number1 := req.Num1
	number2 := req.Num2
	sum := number1 + number2
	res := &calculatorpb.SumResponse{
		Result: sum,
	}
	return res, nil
}

func (*server) CalcPrime(req *calculatorpb.CalcPrimeRequest, stream calculatorpb.CalculateService_CalcPrimeServer) error {
	fmt.Println("Cal-Prime server running", req)
	number := req.GetNumber()
	divisor := int32(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&calculatorpb.CalcPrimeResponse{
				PrimeFactor: divisor,
			})
			number = number / divisor
		} else {
			divisor++
			fmt.Println("Divisor has increased to \n", divisor)
		}

	}
	return nil
}

func main() {
	fmt.Println("hi Cal")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("fail %v", err)

	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculateServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v:", err)
	}

	calculatorpb.RegisterCalculateServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}
}
