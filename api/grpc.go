package api

import (
	context "context"
	pb "godoallserver/proto"
)

type GRPCServer struct {
	pb.UnimplementMyServiceServer
}

func (s *GRPCServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello from gRPC"}, nil
}