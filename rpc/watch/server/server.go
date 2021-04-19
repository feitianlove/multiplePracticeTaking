package main

import (
	"fmt"
	"math/rand"
	"net"
	"net/rpc"
	"sync"
	"time"
)

type KVStoreService struct {
	m      map[string]string
	filter map[string]func(key string)
	mu     sync.Mutex
}

func NewKVStoreService() *KVStoreService {
	return &KVStoreService{
		m:      make(map[string]string),
		filter: make(map[string]func(key string)),
	}
}

func (p *KVStoreService) Get(key string, value *string) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if v, ok := p.m[key]; ok {
		*value = v
		return nil
	}

	return fmt.Errorf("not found")
}

func (p *KVStoreService) Set(kv [2]string, reply *struct{}) error {
	p.mu.Lock()
	defer p.mu.Unlock()

	key, value := kv[0], kv[1]
	fmt.Printf("%+v", p.filter)
	if oldValue := p.m[key]; oldValue != value {
		for _, fn := range p.filter {
			fn(key)
		}
	}

	p.m[key] = value
	return nil
}
func (p *KVStoreService) Watch(timeoutSecond int, keyChanged *string) error {
	//id := fmt.Sprintf("watch-%s-%03d", time.Now(), rand.Int())
	id := fmt.Sprintf("watch-%03d", rand.Int())

	ch := make(chan string, 10) // buffered

	p.mu.Lock()
	p.filter[id] = func(key string) { ch <- key }
	p.mu.Unlock()
	fmt.Println(p.filter)
	select {
	case <-time.After(time.Duration(timeoutSecond) * time.Second):
		return fmt.Errorf("timeout")
	case key := <-ch:
		*keyChanged = key
		return nil
	}
}

func main() {
	// 将KVStroeService对象注册为一个RPC服务
	// 将对象中所有满足RPC规则的对象方法注册为 RPC函数
	// 所有注册方法会放在KVStoreService服务空间执行
	if err := rpc.RegisterName("KVStoreService", NewKVStoreService()); err != nil {
		panic(err)
	}
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	rpc.ServeConn(conn)
}
