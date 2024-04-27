package main

import (
	"log"
	// "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "gRPCgo/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	// SSL certificate
	opts := []grpc.DialOption{}
	tsl := true //change that to false if needed

	if tsl {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("error while loading CA trusted certificate: %v\n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c) // unity
	// doGreetManyTimes(c) // server_streaming
	// doLongGreet(c) // client_streaming
	// doGreetEveryone(c) // bi-directional_streaming
	// doGreetwithdeadline(c, 2*time.Second) // unity but with deadline
}
