package main

import (
	"context"
	"log"

	pb "gRPCgo/calculator/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorServicesClient, n int32){
	log.Println("doSqrt is invoked")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("Error message from server: %s\n", e.Message())
			log.Printf("Error code from server: %s\n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Printf("A negative number was sent!")
				return
			}

		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}

	}

	log.Printf("Sqrt: %f \n", res.Result)
}