package main

import (
	"context"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/publish_subscribe"
	"github.com/moby/moby/pkg/pubsub"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"strings"
	"time"
)

/*
	前面讲述的基于证书的认证是针对每个gRPC链接的认证。
	gRPC还为每个gRPC方法调用提供了认证支持，这样就基于用户Token对不同的方法访问进行权限管理。
*/

type PubsubService struct {
	pub *pubsub.Publisher
}

func NewPubsubService() *PubsubService {
	return &PubsubService{
		pub: pubsub.NewPublisher(100*time.Millisecond, 10),
	}
}

func (p *PubsubService) Publish(
	ctx context.Context, arg *pb.String,
) (*pb.String, error) {
	p.pub.Publish(arg.GetValue())
	return &pb.String{}, nil
}

func (p *PubsubService) Subscribe(
	arg *pb.String, stream pb.PubsubService_SubscribeServer,
) error {
	ch := p.pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, arg.GetValue()) {
				return true
			}
		}
		return false
	})

	for v := range ch {
		if err := stream.Send(&pb.String{Value: v.(string)}); err != nil {
			return err
		}
	}

	return nil
}
func main() {
	// 添加一个证书
	creds, err := credentials.NewServerTLSFromFile("../server.crt", "../server.key")
	if err != nil {
		panic(err)
	}
	//grpc.NewServer()构造一个gRPC服务对象，需要传进去证书
	grpcServer := grpc.NewServer(grpc.Creds(creds))

	pb.RegisterPubsubServiceServer(grpcServer, NewPubsubService())

	conn, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	func() {
		_ = grpcServer.Serve(conn)
	}()
}
