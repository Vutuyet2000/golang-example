package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/blog-grpc/blog/proto"
)

func readBlog(c pb.BlogServiceClient, in string) *pb.Blog {
	log.Println("readBlog was invoked")

	id := &pb.BlogId{
		Id: in,
	}

	res, err := c.ReadBlog(context.Background(), id)
	if err != nil {
		log.Fatalf("Error happened while reading: %v\n", err)
	}
	log.Printf("Blog: %v\n", res)
	return res
}
