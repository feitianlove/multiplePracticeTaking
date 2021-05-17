package main

import (
	"bufio"
	"bytes"
	"container/list"
	"encoding/binary"
	"fmt"
	"golang.org/x/net/html"
	"net"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8848"
)

var mutex sync.Mutex

func main() {
	tcpAddr1, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8848")
	if err != nil {
		panic(err)
	}
	conn1, err := net.DialTCP("tcp", nil, tcpAddr1)
	if err != nil {
		panic(err)
	}

	tcpAddr2, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8849")
	if err != nil {
		panic(err)
	}
	conn2, err := net.DialTCP("tcp", nil, tcpAddr2)
	if err != nil {
		panic(err)
	}

	tcpAddr3, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8850")
	if err != nil {
		panic(err)
	}
	conn3, err := net.DialTCP("tcp", nil, tcpAddr3)
	if err != nil {
		panic(err)
	}
	//写文件
	path := "/Users/ftfeng/goCode/multiplePracticeTaking/distributed/distribute_crawler/email.txt"
	saveFile, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	save := bufio.NewWriter(saveFile)
	defer func() {
		_ = saveFile.Close()
	}()
	var conn = make([]net.Conn, 3)
	go ServerMsgHandler(conn1, save)
	go ServerMsgHandler(conn2, save)
	go ServerMsgHandler(conn3, save)
	conn[0] = conn1
	conn[1] = conn2
	conn[2] = conn3

	//=================================
	urlMap := make(map[string]int)
	count := 0
	go func() {
		url := "http://bbs.tianya.cn/post-140-393973-1.shtml"
		myQueue := list.New()
		myQueue.PushBack(url)
		for myQueue.Len() != 0 {
			myUrl := myQueue.Front()
			// 发送给server
			mystr := myUrl.Value.(string)
			mystrLength := len(mystr)
			mybstart, _ := IntToBytes(mystrLength)
			_, _ = conn[count%len(conn)].Write(mybstart)
			_, _ = conn[count%len(conn)].Write([]byte(mystr))

			myQueue.Remove(myUrl)
			urlList, _ := getUrl(myUrl.Value.(string))
			for _, getUrl := range urlList {
				if _, ok := urlMap[getUrl]; ok {
					urlMap[getUrl]++
				} else {
					myQueue.PushBack(getUrl)
					urlMap[getUrl] = 1
				}
			}
			if count == 100 {
				break
			}
			count++
		}
	}()
	time.Sleep(time.Second * 1000)
}

func ServerMsgHandler(conn net.Conn, save *bufio.Writer) {
	for {
		buf1 := make([]byte, 8)
		n, err := conn.Read(buf1)
		if err != nil || n != 8 {
			fmt.Println("server close")
			return
		}
		length, _ := BytesToInt(buf1)
		buf2 := make([]byte, length)
		n, err = conn.Read(buf2)
		if err != nil || n != length {
			fmt.Println("server close")
			return
		}
		//写入文件
		mutex.Lock()
		_, _ = fmt.Fprintln(save, string(buf2)+conn.RemoteAddr().String()+"\n")
		_ = save.Flush()
		fmt.Println("agent接收到搜索结果", string(buf2), conn.RemoteAddr().String())
		mutex.Unlock()
	}
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

func getUrl(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	doc, err := html.Parse(resp.Body)
	return visit(nil, doc), nil
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				if strings.Contains(a.Val, "javascript") || a.Val == "#fabu_anchor" || a.Val == "#" {
					continue
				}
				if a.Val[:7] != "http:" {
					a.Val = "http://www.tianya.cn/519255" + a.Val
				}
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
