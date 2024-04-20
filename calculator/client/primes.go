package main

import (
	"context"
	pb "gRPCgo/calculator/proto"
	// "io"
	"log"
)

func doPrimes(c pb.CalculatorServicesClient) {
	log.Println("doPrimes was invoked")

	req := &pb.PrimeRequest{
		Number: 103,
	}

	stream, err := c.Primes(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling primes: %v \n", err)
	}

	for {
		res, err := stream.Recv()

		log.Printf("Primes: %d \n", res.Result)

		// if err != io.EOF {
		// 	break
		// }

		if err != nil {
			log.Fatalf("error while reading stream: %v \n", err)
		}

		log.Printf("Primes: %d \n", res.Result)
	}
}