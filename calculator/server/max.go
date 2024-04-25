package main

import (
	pb "gRPCgo/calculator/proto"
	"log"
	"io"
)

func (s *Server) Max(stream pb.CalculatorServices_MaxServer) error {
	log.Printf("Max func was invoked")
	var maximum int32  = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		if number := req.Number; number > maximum {
			maximum = number
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("Error while sending data to client: %v\n", err)
			}
		}
	}
}