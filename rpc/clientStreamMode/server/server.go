package main

import (
	"fmt"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/clientStreamMode/pb/clientMode"
	"google.golang.org/grpc"
	"io"
	"net"
)

type server struct {
}

func (s server) ListValue(svr pb.StreamClient_ListValueServer) error {
	for {
		res, err := svr.Recv()
		if err == io.EOF {
			return svr.SendAndClose(&pb.SimpleResponse{Value: "ok"})

		}
		if err != nil {
			return err
		}
		fmt.Printf("server recieve %s\n", res)
	}
}

//=======
const (
	Address = ":8080"
	Network = "tcp"
)

func main() {
	conn, err := net.Listen(Network, Address)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStreamClientServer(grpcServer, &server{})
	err = grpcServer.Serve(conn)
	if err != nil {
		panic(err)
	}
}
