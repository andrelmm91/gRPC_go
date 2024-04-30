package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "gRPCgo/blog/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
	}
	defer conn.Close()

	c := pb.NewBlogServiceClient(conn)

	// // mocking creating a Blog
	id := createBlog(c)
	// // mocking reading a Blog
	readBlog(c, id) //valid
	// readBlog(c, "aNonExistingId") //invalid
	// //mocking updating a Blog
	updateBlog(c, id)
	// //mocking listing all blogs
	listBlog(c)
}
