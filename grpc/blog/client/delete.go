package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/blog-grpc/blog/proto"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("---deleteBlog was invoked---")
	log.Printf("deleteeId: %s\n", id)

	deleteId := &pb.BlogId{
		Id: id,
	}

	_, err := c.DeleteBlog(context.Background(), deleteId)
	if err != nil {
		log.Printf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was deleted")
}
