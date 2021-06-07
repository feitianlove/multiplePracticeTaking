package mapReduce

import (
	"fmt"
	"net"
	"net/rpc"
	"sync"
)

//工作者，等待DOTask任务或者shutDown任务
type Worker struct {
	sync.Mutex
	Name       string // 名称
	Map        func(string, string) []KeyValue
	Reduce     func(string, []string) string
	NTask      int // 执行的任务总量
	ConCurrent int // 当前worker执行的总任务量
	L          net.Listener
	NRpc       int //退出标志
}

/*
	任务执行函数， 再有新任务分配给改worker时，master调用该函数执行任务
*/
func (wk *Worker) DoTask(args *DoTaskArgs, _ *struct{}) error {
	//任务数量+1
	wk.Lock()
	wk.NTask++
	switch args.Phase {
	case mapPhase:
		DoMpa(args.JobName, args.TaskNumer, args.File, args.NumOtherPhase, wk.Map)
	case ReducePhase:
		DoReduce(args.JobName, args.TaskNumer, mergeName(args.JobName, args.TaskNumer), args.NumOtherPhase, wk.Reduce)
	}
	fmt.Printf("%s:%v task #%d is done\n", wk.Name, args.Phase, args.TaskNumer)
	return nil
}

//启动worker，与master建立链接
func RunWorker(masterAddress string, me string, nRpc int,
	mapFun func(string, string) []KeyValue,
	ReduceFunc func(string, []string) string) {
	fmt.Println("run worker")
	wk := new(Worker)
	wk.Name = me
	wk.Map = MapFunc
	wk.Reduce = ReduceFunc
	wk.NRpc = nRpc
	//新建一个rpc服务
	rpcs := rpc.NewServer()
	rpcs.Register(wk)
	//os.Remove(me)
	l, err := net.Listen("unix", me)
	if err != nil {
		fmt.Println("run worker error ", err, me)
		panic(err)
	}
	wk.L = l
	//注册到master
	wk.RegisterWorker(masterAddress)

	for {
		wk.Lock()
		//没有链接上
		if wk.NRpc == 0 {
			wk.Unlock()
			break
		}
		conn, err := wk.L.Accept()
		if nil != err {
			wk.Unlock()
			break
		} else {
			wk.Unlock()
			go rpcs.ServeConn(conn)
		}
		wk.L.Close()
		fmt.Printf("RunWorker%s is exit\n", me)
	}
}

//告知master worker的存在
func (wk *Worker) RegisterWorker(master string) {
	args := new(RegisterArgs)
	args.Worker = wk.Name
	//调用master的注册函数，注册worker
	ok := call(master, "Master.Register", args, new(struct{}))
	if !ok {
		fmt.Printf("Register: RPC %s worker err \n", master)
	}
}

//紧急中断
func (wk *Worker) Close(_ *struct{}, res *ShutDownReplay) error {
	fmt.Println("shutdown register server\n", wk.Name)
	wk.Lock()
	wk.NRpc = 1
	defer wk.Unlock()
	res.NTasks = wk.NTask
	return nil
}
