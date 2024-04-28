package main

import (
	"context"
	"log"

	pb "gRPCgo/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---CreateBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "andre",
		Title:    "my first blog",
		Content:  "content123",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
