package main

import (
	"context"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/token"
	"google.golang.org/grpc"
	"log"
)

type Authentication struct {
	User     string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"user": a.User, "password": a.Password}, nil
}
func (a *Authentication) RequireTransportSecurity() bool {
	return false
}

var (
	port = ":5000"
)

func main() {
	auth := Authentication{
		User:     "gopher",
		Password: "password",
	}
	conn, err := grpc.Dial("localhost"+port, grpc.WithInsecure(), grpc.WithPerRPCCredentials(&auth))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: "gopher"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("doClientWork: %s", r.Message)
}
