package main

import (
	"context"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/expand_grpc"
	"google.golang.org/grpc"
	"net"
)

type RestServiceImpl struct{}

func (r *RestServiceImpl) Get(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "Get hi:" + message.Value + "#"}, nil
}

func (r *RestServiceImpl) Post(ctx context.Context, message *pb.StringMessage) (*pb.StringMessage, error) {
	return &pb.StringMessage{Value: "Post hi:" + message.Value + "@"}, nil
}
func main() {
	grpcServer := grpc.NewServer()
	pb.RegisterRestServiceServer(grpcServer, new(RestServiceImpl))
	lis, _ := net.Listen("tcp", ":5000")
	defer func() {
		_ = grpcServer.Serve(lis)
	}()
}

// curl localhost:8080/get/gopher
// curl localhost:8080/post -X POST --data '{"value":"grpc"}'
