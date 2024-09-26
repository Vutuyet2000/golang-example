package main

import (
	pb "github.com/Clement-Jean/blog-grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	Title    string             `bson:"title"`
	AuthorID string             `bson:"author_id"`
	Content  string             `bson:"content"`
}

type CreateBlogItem struct {
	Title    string `bson:"title"`
	AuthorID string `bson:"author_id"`
	Content  string `bson:"content"`
}

func documentToBlog(data *BlogItem) *pb.Blog {
	return &pb.Blog{
		Id:       data.ID.Hex(),
		Title:    data.Title,
		AuthorId: data.AuthorID,
		Content:  data.Content,
	}
}
