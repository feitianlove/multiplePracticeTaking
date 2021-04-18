package main

import (
	"fmt"
	pb "github.com/feitianlove/Miscellaneous.git/rpc/bothStreamMode/pb"
	"google.golang.org/grpc"
	"io"
	"net"
	"strconv"
	"time"
)

type server struct {
}

func (s server) Conversations(svr pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := svr.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = svr.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		fmt.Printf("SERVER RECV%s\n", req)
		time.Sleep(1 * time.Second)
	}
}

const (
	Network = "tcp"
	Address = ":8080"
)

func main() {
	conn, err := net.Listen(Network, Address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamServer(grpcServer, &server{})

	err = grpcServer.Serve(conn)
	if err != nil {
		panic(err)
	}

}
