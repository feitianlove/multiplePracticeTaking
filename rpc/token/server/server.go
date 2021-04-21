package main

import (
	"fmt"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/token"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log"
	"net"
)

var (
	port = ":5000"
)

type myGrpcServer struct{}

// https://chai2010.cn/advanced-go-programming-book/ch4-rpc/ch4-05-grpc-hack.html 也可以参考实现auth方法，这里仅仅是个例子

func (s *myGrpcServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing credentials")
	}

	var (
		user   string
		passwd string
	)

	if val, ok := md["user"]; ok {
		user = val[0]
	}
	if val, ok := md["password"]; ok {
		passwd = val[0]
	}

	if user != "gopher" || passwd != "password" {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid token: appid=%s, appkey=%s", user, passwd)
	}

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// 实现grpc.PerRPCCredentials接⼝

func main() {
	startServer()

}

func startServer() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}
