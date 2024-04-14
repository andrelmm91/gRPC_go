package main

import (
	"log"
	"net"

	pb "gRPCgo/calculator/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.CalculatorServicesServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
	}

	log.Printf("listening on %s \n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServicesServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}