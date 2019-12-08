package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bri997/grpc-go-course/greet/hands-on/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func main() {

	fmt.Println("hi")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("fail %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreatServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server %v", err)
	}

}
