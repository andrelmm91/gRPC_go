package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gRPCgo/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServicesClient(conn)

	// doSum(c)
	// doPrimes(c)
	// doAvg(c)
	doMax(c)
}