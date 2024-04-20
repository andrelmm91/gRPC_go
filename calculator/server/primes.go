package main

import (
	pb "gRPCgo/calculator/proto"
	"log"
)

func (s *Server) Primes(in *pb.PrimeRequest, stream pb.CalculatorServices_PrimesServer) error {
	log.Printf("primes found was involked with %v \n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number % divisor == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: divisor,
			})

			number /= divisor
		} else {
			divisor++
		}
	}

	return nil
}