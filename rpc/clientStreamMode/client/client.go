package main

import (
	"context"
	"fmt"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/clientStreamMode/pb/clientMode"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

const Address string = ":8080"

func ListValue() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	defer func() {
		_ = conn.Close()
	}()
	if err != nil {
		panic(err)
	}
	streamClient := pb.NewStreamClientClient(conn)
	stream, err := streamClient.ListValue(context.Background())
	if err != nil {
		panic(err)
	}
	for n := 0; n < 5; n++ {
		//向流中发送消息
		err := stream.Send(&pb.StreamRequest{Data: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			fmt.Printf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	fmt.Printf("client send success%s\n", res)

}
func main() {
	ListValue()
}
