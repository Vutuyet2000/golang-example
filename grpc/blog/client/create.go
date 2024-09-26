package main

import (
	"context"
	"log"

	pb "github.com/Clement-Jean/blog-grpc/blog/proto"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Printf("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Snow",
		Title:    "This is my blog",
		Content:  "This is the content of this blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %v\n", res)
	return res.Id
}
