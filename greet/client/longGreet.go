package main

import (
	"context"
	pb "gRPCgo/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient){
	log.Println("doLongGreet is invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "andre1"},
		{FirstName: "andre2"},
		{FirstName: "andre3"},
		{FirstName: "andre4"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("could not LongGreet: %v \n", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from LongGreet: %v \n", err)
	}

	log.Printf("LongGreet: %s \n", res.Result)
}