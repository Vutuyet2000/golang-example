package main

import pb "github.com/Clement-Jean/blog-grpc/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
