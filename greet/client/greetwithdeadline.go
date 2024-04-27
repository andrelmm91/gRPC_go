package main

import (
	"context"
	"log"
	"time"

	pb "gRPCgo/greet/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetwithdeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetwithdeadline is invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	res, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "andre",
	})

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceed")
				return
			} else {
				log.Fatalf("Unexpected gRPC erre %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC errpr: %v\n", err)
		}
	}

	log.Printf("doGreetwithdeadline: %s \n", res.Result)
}
