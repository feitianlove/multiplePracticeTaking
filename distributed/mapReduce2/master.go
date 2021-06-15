package mapReduce

import (
	"fmt"
	"net"
	"sync"
)

type Master struct {
	//master ip
	Address string
	//锁
	sync.Mutex
	//存储worker的缓存、存储套接字，代表rpc地址
	Workers []string
	//jobName
	JobName string
	//
	NewCond *sync.Cond
	//输入文件
	Files []string
	//分区数量
	NReduce int
	//RPC Listener
	L net.Listener
	//中断信号
	ShutDown chan struct{}
	Status   []int
}

func NewMaster(address string) *Master {
	mr := &Master{
		Address:  address,
		ShutDown: make(chan struct{}),
	}
	mr.NewCond = sync.NewCond(mr)
	return mr
}

//实现一个worker函数，这是一哥rpc方法
//2021/06/08 15:09:35 rpc.Register: method "Register" has 4 input parameters; needs exactly three
//func (mr *Master) Register(_, _ *struct{}, args *RegisterArgs) error {
func (mr *Master) Register(args *RegisterArgs, _ *struct{}) error {
	mr.Lock()
	defer mr.Unlock()
	mr.Workers = append(mr.Workers, args.Worker)
	//广播给所有
	fmt.Println("mr.workers 队列 ", args, mr.Workers)
	mr.NewCond.Broadcast()
	return nil
}

// 实现一个worker传递函数，将所以已存在的worker与新注册的worker传递到一个通道中，让调度函数进行处理
func (mr *Master) ForwardRegister(ch chan string) {

	i := 0
	for {
		mr.Lock()
		if len(mr.Workers) > i {
			w := mr.Workers[i]
			go func() {
				ch <- w
			}()
			i++
		} else {
			fmt.Println("ForwardRegister", mr.Workers, i)
			mr.NewCond.Wait()
		}
		mr.Unlock()

	}

}

//分布式执行mapReduce任务,通过rpc在主服务器上注册的workers调度map和reduce任务
func Distributed(jobName string, files []string, nReduce int, master string) {
	fmt.Println("distributed....")
	var wg sync.WaitGroup
	mr := NewMaster(master)
	//启动master的RPC服务
	wg.Add(1)
	go mr.StartRpcServer(&wg)
	//执行任务
	fmt.Println("Distributed 中的mr", mr)
	go mr.Run(jobName, files, nReduce, func(phase JobPhase) {
		fmt.Println("run")
		ch := make(chan string)
		go mr.ForwardRegister(ch)
		//调度执行
		schedule(mr.JobName, mr.Files, mr.NReduce, phase, ch)
	}, func() {
		mr.Status = mr.KillWokre()
		mr.StopRpcServer()
	})
	wg.Wait()
}
func (mr *Master) Run(jobName string, files []string, nReduce int,
	schedule func(phase JobPhase),
	finish func()) { //finsh收尾
	mr.JobName = jobName
	mr.NReduce = nReduce
	mr.Files = files
	//执行map任务
	schedule(mapPhase)
	fmt.Println("开始执行reduce")
	//执行reduce热卖
	schedule(ReducePhase)
	finish()
	//合并文件
	mr.merge()
}

//清理worker
//该函数通过向每个worker发送shutDown RPC 请求来清理worker
//返回worker已经执行的任务数量
func (mr *Master) KillWokre() []int {
	mr.Lock()
	nTasks := make([]int, len(mr.Workers))
	defer mr.Unlock()
	for _, w := range mr.Workers {
		fmt.Printf("master shutdown worker %s\n", w)
		var replay ShutDownReplay
		//调用worker的shutdown方法
		ok := call(w, "Worker.Close", new(struct{}), &replay)
		if !ok {
			fmt.Printf("Master: shutdown worker %s failed \n", w)
		} else {
			nTasks = append(nTasks, replay.NTasks)
		}

	}
	return nTasks
}
