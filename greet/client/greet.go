package main

import(
	"context"
	"log"

	pb "gRPCgo/greet/proto"
)

func doGreet(c pb.GreetServiceClient){
	log.Println("doGreet is invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "andre",
	})

	if err != nil {
		log.Fatalf("could not greet: %v \n", err)
	}

	log.Printf("greeting: %s \n", res.Result)
}