package mapReduce

import (
	"fmt"
	"net"
	"net/rpc"
	"os"
	"sync"
)

/*
	master与worker进行通信互联、worker注册
	实现rpc功能： 开启rcp等待worker注册、中断master RPC服务、正常停止rpc服务

*/

type RegisterArgs struct {
	Worker string
}

func (mr *Master) StartRpcServer(wg *sync.WaitGroup) {
	//新建一个rpc实例对象
	rpcs := rpc.NewServer()
	//注册master方法
	err := rpcs.Register(mr)
	os.Remove(mr.Address)
	if err != nil {
		fmt.Println("init RPC err", err)
		panic(err)
	}
	//监听

	listen, err := net.Listen("unix", mr.Address)
	if err != nil {
		fmt.Println("监听rpc unix err ", err)
		panic(err)
	}
	fmt.Println("master 监听rpc unix successs", mr.Address)
	mr.L = listen
	//监听地址，获取链接
	go func() {
		fmt.Println("master 等待worker中")
	loop:
		for {
			//判断是否shudown
			select {
			case <-mr.ShutDown:
				break loop
			default:

			}
			conn, err := listen.Accept()
			if err != nil {
				fmt.Println("Register server err, please try again ", err)
				break
			} else {
				go func() {
					rpcs.ServeConn(conn)
					conn.Close()
				}()
			}
		}
		wg.Done()
		fmt.Println("registerServer done")
	}()
}

//紧急中断
func (mr *Master) Close(_, _ *struct{}) error {
	fmt.Println("shutdown register server")
	close(mr.ShutDown)
	mr.L.Close()
	return nil
}

//正常中断
func (mr *Master) StopRpcServer() {
	var reply ShutDownReplay
	ok := call(mr.Address, "Master.Close", new(struct{}), &reply)
	fmt.Println("StopRpcServer", ok)
}
