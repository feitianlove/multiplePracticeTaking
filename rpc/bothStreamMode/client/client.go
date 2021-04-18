package main

import (
	"context"
	"fmt"
	"github.com/feitianlove/Miscellaneous.git/rpc/bothStreamMode/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"strconv"
)

const Address = ":8080"

func Conversations() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	streamClient := pb.NewStreamClient(conn)
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		panic(err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		fmt.Printf("client recieve %s\n", res)
	}
}

func main() {
	Conversations()
}
