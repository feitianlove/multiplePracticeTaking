package main

import (
	"fmt"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/serverStreamMode/pb/clientMode"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type Server struct {
}

func (s *Server) ListValue(request *pb.SimpleRequest, svr pb.StreamServer_ListValueServer) error {
	fmt.Printf("server request :%s\n", request)
	for n := 0; n < 5; n++ {
		fmt.Println("----------------")
		err := svr.Send(&pb.StreamResponse{
			StreamValue: request.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	// 默认单次接收最大消息长度为`1024*1024*4`bytes(4M)，单次发送消息最大长度为`math.MaxInt32`bytes
	// grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*4), grpc.MaxSendMsgSize(math.MaxInt32))
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterStreamServerServer(grpcServer, &Server{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
