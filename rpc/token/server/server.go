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
	server := grpc.NewServer(grpc.UnaryInterceptor(filter))
	pb.RegisterGreeterServer(server, new(myGrpcServer))

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Panicf("could not list on %s: %s", port, err)
	}

	if err := server.Serve(lis); err != nil {
		log.Panicf("grpc serve error: %s", err)
	}
}

// 过滤器
/*
	gRPC中的grpc.UnaryInterceptor和grpc.StreamInterceptor分别对普通方法和流方法提供了截取器的支持。我们这里简单介绍普通方法的截取器用法。
	要实现普通方法的截取器，需要为grpc.UnaryInterceptor的参数实现一个函数：
	注意：
	不过gRPC框架中只能为每个服务设置一个截取器，因此所有的截取工作只能在一个函数中完成。开源的grpc-ecosystem项目中的go-grpc-middleware包已经
	基于gRPC对截取器实现了链式截取器的支持。
*/
func filter(ctx context.Context,
	req interface{}, info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Println("fileter:", info)
	log.Println("req:", req)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	return handler(ctx, req)
}
