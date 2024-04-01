package main

import (
    pb "github.com/ikatolkar/grpc-server-client/proto"
    "context"

)


func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
    return &pb.HelloResponse{
        Message: "Hello",
    }, nil
}
