package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"sync"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8849"
)

// 解决线程安全
var mutex sync.Mutex

var reg = "[A-Za-z0-9]+@[a-zA-Z0-9_-]+(.[a-zA-Z0-9_-]+)+"

func getEmail(url string, conn net.Conn) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	rex := regexp.MustCompile(reg)

	res := rex.FindAllString(string(data), -1)
	mutex.Lock()
	for _, email := range res {
		if len(email) != 0 {
			mystr := email
			mystrLength := len(mystr)
			mybstart, _ := IntToBytes(mystrLength)
			_, _ = conn.Write(mybstart)
			_, _ = conn.Write([]byte(mystr))
		}
	}
	mutex.Unlock()
	return nil
}

//func Worker(id int, jobs <-chan string, result chan<- string) {
//	url := <-jobs
//	templist, _ := getEmail(url)
//	for _, email := range templist {
//		result <- email
//		fmt.Println("<-", email, id)
//	}
//}
//func main1111() {
//	jobs := make(chan string, 100)   // 任务url
//	result := make(chan string, 100) //结构result
//	var count int
//	// 去重url
//	urlMap := make(map[string]int)
//	go func() {
//		url := "http://bbs.tianya.cn/post-140-393973-1.shtml"
//		myQueue := list.New()
//		myQueue.PushBack(url)
//		for myQueue.Len() != 0 {
//			myUrl := myQueue.Front()
//			//fmt.Println(myUrl.Value.(string))
//			jobs <- myUrl.Value.(string)
//			go Worker(count, jobs, result)
//			myQueue.Remove(myUrl)
//			urlList, _ := getUrl(myUrl.Value.(string))
//			for _, getUrl := range urlList {
//				if _, ok := urlMap[getUrl]; ok {
//					urlMap[getUrl]++
//				} else {
//					myQueue.PushBack(getUrl)
//					urlMap[getUrl] = 1
//				}
//			}
//			if count == 100 {
//				break
//			}
//			count++
//		}
//
//	}()
//	for {
//		fmt.Println("get<-", <-result)
//	}
//}

func Server() {
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = listener.Close()
	}()
	fmt.Println("服务器开启", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("server", err)

		}
		fmt.Println("server", "有客户加入", conn.RemoteAddr())

		go msgHandler(conn)
	}
}

func msgHandler(conn net.Conn) {
	for {
		buf1 := make([]byte, 8)
		n, err := conn.Read(buf1)
		if err != nil || n != 8 {
			fmt.Println("client close")
			return
		}
		length, _ := BytesToInt(buf1)
		buf2 := make([]byte, length)
		n, err = conn.Read(buf2)
		if err != nil || n != length {
			fmt.Println("client close")
			return
		}
		fmt.Println("server收到，即将爬取结果", string(buf2))
		url := string(buf2)
		go func() {
			_ = getEmail(url, conn) //读到一个url
		}()
	}
}

func main() {
	Server()
}

func BytesToInt(bts []byte) (int, error) {
	byteBuf := bytes.NewBuffer(bts)
	var data int64
	err := binary.Read(byteBuf, binary.BigEndian, &data)
	if err != nil {
		return 0, err
	}
	return int(data), nil
}
func IntToBytes(n int) ([]byte, error) {
	data := int64(n)
	byteBuf := bytes.NewBuffer([]byte{})
	err := binary.Write(byteBuf, binary.BigEndian, data)
	if err != nil {
		return nil, err
	}
	return byteBuf.Bytes(), nil
}
