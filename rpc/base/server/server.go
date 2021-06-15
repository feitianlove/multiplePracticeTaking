package main

import (
	"log"
	"net"
	"net/rpc"
)

// 构造一个hello service类型
type HelloService struct {
}

func (p *HelloService) Hello(request string, reply *string) error {

	*reply = "hello:" + request
	return nil
}

func main() {
	err := rpc.RegisterName("feitian", new(HelloService))
	//err  := rpc.Register(new(HelloService)) // 不用指定默认就是HelloService
	if err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)
	conn, err = listener.Accept()

	rpc.ServeConn(conn)

	// 自定义编码和解码 通过json， 一般原生RPC跨语言的时候使用
	//go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	//client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

}
