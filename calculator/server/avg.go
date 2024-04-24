package main

import (
	pb "gRPCgo/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.CalculatorServices_AvgServer) error {
	log.Printf("Avg function was invoked")

	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		log.Printf("Receiving number: %d \n", req.Number)
		sum += req.Number
		count++
	}
}