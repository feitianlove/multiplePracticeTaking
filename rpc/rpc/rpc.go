package rpc

import (
	"fmt"
	"net"
	"reflect"
)

/*
	维护函数map,客户端传来的东西进行解析, 函数的返回值打包，传给客户端
*/

type Server struct {
	// 地址
	addr string
	// map 用于维护关系的
	funcs map[string]reflect.Value
}

// 构造方法
func NewServer(addr string) *Server {
	return &Server{addr: addr, funcs: make(map[string]reflect.Value)}
}

func (s *Server) Register(rpcName string, f interface{}) {
	if _, ok := s.funcs[rpcName]; ok {
		return
	}
	fVal := reflect.ValueOf(f)
	s.funcs[rpcName] = fVal
}

func (s *Server) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("监听 %s err :%v", s.addr, err)
	}
	for {
		conn, err := lis.Accept()
		if err != nil {
			return err
		}
		serSession := NewSession(conn)
		b, err := serSession.Read()
		if err != nil {
			return err
		}
		// 数据解码
		rpcData, err := decode(b)
		if err != nil {
			return err
		}
		fmt.Println(rpcData)
		// 根据读到的name，得到要调用的函数
		f, ok := s.funcs[rpcData.Name]
		if !ok {
			return fmt.Errorf("函数 %s 不存在", rpcData.Name)
		}
		// 遍历解析客户端传来的参数,放切片里
		inArgs := make([]reflect.Value, 0, len(rpcData.Args))
		for _, arg := range rpcData.Args {
			inArgs = append(inArgs, reflect.ValueOf(arg))
		}
		// 反射调用方法
		// 返回Value类型，用于给客户端传递返回结果,out是所有的返回结果
		out := f.Call(inArgs)
		// 遍历out ，用于返回给客户端，存到一个切片里
		outArgs := make([]interface{}, 0, len(out))
		for _, o := range out {
			outArgs = append(outArgs, o.Interface())
		}
		// 数据编码，返回给客户端
		respRPCData := RPCData{rpcData.Name, outArgs}
		bytes, err := encode(respRPCData)
		if err != nil {
			return err
		}
		// 将服务端编码后的数据，写出到客户端
		err = serSession.Write(bytes)
		if err != nil {
			return err
		}
	}
}

/*====================客户端=========================*/
type Client struct {
	conn net.Conn
}

// 构造方法
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// 实现通用的RPC客户端
// 传入访问的函数名
// fPtr指向的是函数原型
//var select fun xx(User)
//cli.callRPC("selectUser",&select)

func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	fn := reflect.ValueOf(fPtr).Elem()
	f := func(args []reflect.Value) []reflect.Value {
		// 处理参数
		inArgs := make([]interface{}, 0, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		// 连接
		cliSession := NewSession(c.conn)
		// 编码数据
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}
		// 写数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		// 服务端发过来返回值，此时应该读取和解析
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		// 解码
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		// 处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			// 必须进行nil转换
			if arg == nil {
				// reflect.Zero()会返回类型的零值的value
				// .out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}
	// 完成原型到函数调用的内部转换
	// 参数1是reflect.Type
	// 参数2 f是函数类型，是对于参数1 fn函数的操作
	// fn是定义，f是具体操作
	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数fPtr赋值，过程
	fn.Set(v)

}
