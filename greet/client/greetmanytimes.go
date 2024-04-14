package main

import (
	"context"
	"io"
	"log"

	pb "gRPCgo/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient){
	log.Println("doGreetManyTimes is invoked")

	req := &pb.GreetRequest{
		FirstName: "andre",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("could not GreetManyTimes: %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatalf("error while reading the stream: %v \n", err)
		}

		log.Printf("Greeting Many Times: %s \n", msg.Result)
	}
}