package main

import (
	"context"
	s "github.com/feitianlove/multiplePracticeTaking/rpc/rpc_x/server"
	"github.com/smallnest/rpcx/client"
	"log"
	"time"
)

func main1() {
	//#1 定义了使用什么方式来实现服务发现。 在这里我们使用最简单的 Peer2PeerDiscovery（点对点）。客户端直连服务器来获取服务地址。
	d, err := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	if err != nil {
		panic(err)
	}
	// 创建了 XClient, 并且传进去了 FailMode、 SelectMode 和默认选项。 FailMode 告诉客户端如何处理调用失败：重试、快速返回，或者 尝试另一
	//台服务器。 SelectMode 告诉客户端如何在有多台服务器提供了同一服务的情况下选择服务器。
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	args := &s.Calculate{
		A: 10,
		B: 20,
	}
	r := &s.Result{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	err = xclient.Call(ctx, "Mul", args, r)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	cancel()
	log.Printf("%d * %d = %d", args.A, args.B, r.R)
}

// 异步调用
func main() {
	//#1 定义了使用什么方式来实现服务发现。 在这里我们使用最简单的 Peer2PeerDiscovery（点对点）。客户端直连服务器来获取服务地址。
	d, err := client.NewPeer2PeerDiscovery("tcp@127.0.0.1:8972", "")
	if err != nil {
		panic(err)
	}
	// 创建了 XClient, 并且传进去了 FailMode、 SelectMode 和默认选项。 FailMode 告诉客户端如何处理调用失败：重试、快速返回，或者 尝试另一
	//台服务器。 SelectMode 告诉客户端如何在有多台服务器提供了同一服务的情况下选择服务器。
	xclient := client.NewXClient("Arith", client.Failtry, client.RandomSelect, d, client.DefaultOption)

	args := &s.Calculate{
		A: 10,
		B: 20,
	}
	r := &s.Result{}

	call, err := xclient.Go(context.Background(), "Mul", args, r, nil)
	if err != nil {
		log.Fatalf("failed to call: %v", err)
	}
	replyCall := <-call.Done
	if replyCall.Error != nil {
		log.Fatalf("failed to call: %v", replyCall.Error)
	} else {
		log.Printf("%d * %d = %d", args.A, args.B, r.R)
	}
}
