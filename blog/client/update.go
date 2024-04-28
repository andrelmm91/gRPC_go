package main

import (
	"context"
	pb "gRPCgo/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("---UpdateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "not Andre",
		Title:    "new title",
		Content:  "content with adds",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("Unexpected error while updating: %v\n", err)
	}

	log.Printf("Blog was updated")
}
