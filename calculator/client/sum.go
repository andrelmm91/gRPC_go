package main

import (
	"context"
	"log"

	pb "gRPCgo/calculator/proto"
)

func doSum(c pb.CalculatorServicesClient){
	log.Println("doSum is invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber: 1,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatalf("could not greet: %v \n", err)
	}

	log.Printf("greeting: %d \n", res.Result)
}