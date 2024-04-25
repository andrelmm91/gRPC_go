package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "gRPCgo/calculator/proto"
)

func doMax(c pb.CalculatorServicesClient){
	log.Println("doMax is invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("could not GreetEveryone: %v \n", err)
	}

	waitc := make(chan struct{})

	//sending
	go func() {
		numbers := []int32{4, 7, 2, 19, 4, 6, 32}

		for  _, number := range numbers {
			log.Printf("Send request: %v \n", number)
			stream.Send(&pb.MaxRequest{
				Number: number,
			})

			time.Sleep(500 * time.Millisecond)
		}

		stream.CloseSend()
	}()
	
	//receiving
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

			log.Printf("Received a new maximum: %d\n", res.Result)
		}

		close(waitc)
	}()

	// will wait those two go routines send and receive and at the end close(waitc) is executed and go out of scope and returns
	<-waitc
}