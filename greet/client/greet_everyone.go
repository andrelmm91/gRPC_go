package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "gRPCgo/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient){
	log.Println("GreetEveryone is invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("could not GreetEveryone: %v \n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "andre1"},
		{FirstName: "andre2"},
		{FirstName: "andre3"},
		{FirstName: "andre4"},
	}

	waitc := make(chan struct{})

	go func() {
		for  _, req := range reqs {
			log.Printf("Send request: %v \n", req)
			stream.Send(req)

			time.Sleep(100 * time.Millisecond)
		}

		stream.CloseSend()
	}()

	go func() {
		for {
			res, err :=  stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Println("Error while receiving %v\n", err)
				break
			}

			log.Printf("Received: %v\n", res.Result)
		}

		close(waitc)
	}()

	// will wait those two go routines send and receive and at the end close(waitc) is executed and go out of scope and returns
	<-waitc
}