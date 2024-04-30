package main

import (
	"context"
	pb "gRPCgo/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---ReadBlog was invoked")

	req := &pb.BlogId{Id: id}
	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("Unexpected error while reading: %v\n", err)
	}

	log.Printf("Blog has been read: %s\n", res)
	return res
}
