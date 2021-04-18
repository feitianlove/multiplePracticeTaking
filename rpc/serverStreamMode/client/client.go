package main

import (
	"context"
	"fmt"
	pb "github.com/feitianlove/Miscellaneous.git/rpc/serverStreamMode/pb/clientMode"
	"google.golang.org/grpc"
	"io"
	"log"
)

const (
	// Address 连接地址
	Address string = ":8000"
)

func client() {

	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	defer func() {
		_ = conn.Close()
	}()
	req := pb.SimpleRequest{
		Data: "I'm ftfeng, this is server stream mode",
	}
	grpcClient := pb.NewStreamServerClient(conn)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		panic(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		fmt.Println(res)
	}

}

func main() {
	client()
}
