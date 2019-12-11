package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/bri997/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hi Client...")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)
	// fmt.Printf("Created client: %f", c)

	// Change function to do different API
	// doUnary(c)
	doServerStreaming(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do a Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Brian",
			LastName:  "Musial",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		println("req =", req)
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("starting to do Server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Brian",
			LastName:  "Musial",
		},
	}

	resStream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling GreetManyTimes RPC %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			//We have reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v, err")
		}
		log.Printf("Response for GreetManyTimes: %v", msg.GetResult())
	}

}
