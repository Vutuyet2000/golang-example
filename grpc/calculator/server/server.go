package main

import pb "github.com/Clement-Jean/grpc/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
