package main

import (
	"context"
	pb "github.com/feitianlove/multiplePracticeTaking/rpc/publish_subscribe"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

// grpc.WithInsecure()选项跳过了对服务器证书的验证

func main() {
	// 在客户端基于服务器的证书和服务器名字就可以对服务器进行验证
	// 表示服务器的名字为server.grpc.io
	//其中redentials.NewClientTLSFromFile是构造客户端用的证书对象，第一个参数是服务器的证书文件，第二个参数是签发证书的服务器的名字。
	//然后通过grpc.WithTransportCredentials(creds)将证书对象转为参数选项传人grpc.Dial函数。
	creds, err := credentials.NewClientTLSFromFile(
		"server.crt", "server.grpc.io",
	)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial("localhost:1234", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	client := pb.NewPubsubServiceClient(conn)

	_, err = client.Publish(
		context.Background(), &pb.String{Value: "golang: hello Go"},
	)
	if err != nil {
		log.Fatal(err)
	}
	_, err = client.Publish(
		context.Background(), &pb.String{Value: "docker: hello Docker"},
	)
	if err != nil {
		log.Fatal(err)
	}
}
