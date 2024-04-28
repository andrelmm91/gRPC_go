package main

import (
	"context"
	"log"
	"net"

	pb "gRPCgo/blog/proto"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var collection *mongo.Collection
var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BlogServiceServer
}

func main() {
	// connecting to mongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("failed to connect to Mongo: %v \n", err)
	}

	_, err = mongo.Connect(context.Background())
	if err != nil {
		log.Fatalf("failed to connect to Mongo: %v \n", err)
	}

	collection = client.Database("blogDB").Collection("blog")

	// listening to gRPC
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
	}

	log.Printf("listening on %s \n", addr)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
