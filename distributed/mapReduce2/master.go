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

/*
	任务调度函数(顺序执行的)
	files 要处理的文件
	nReduce 分区数量
*/
func SequentialOne(jobName string, files []string, nReduce int,
	mapF func(string, string) []KeyValue,
	reduceF func(string, []string) string) {

	m := NewMaster("master")
	m.Run(jobName, files, nReduce, func(phase JobPhase) {
		switch phase {
		case mapPhase:
			for i, f := range files {
				DoMpa(m.JobName, i, f, m.NReduce, mapF)
			}
		case ReducePhase:
			for i := 0; i < m.NReduce; i++ {
				DoReduce(m.JobName, i, mergeName(m.JobName, i), len(files), reduceF)
			}
		}
	}, func() {
		m.Status = m.KillWokre()
	})
}

func (m *Master) Run(jobName string, files []string, nReduce int,
	schedule func(phase JobPhase),
	finish func()) { //finsh收尾
	m.JobName = jobName
	m.NReduce = nReduce
	m.Files = files
	//执行map任务
	schedule(mapPhase)
	//执行reduce热卖
	schedule(ReducePhase)
	finish()
	//合并文件
	m.merge()
}

//实现一个worker函数，这是一哥rpc方法
func (mr *Master) Register(args *RegisterArgs) error {
	mr.Lock()
	defer mr.Unlock()
	mr.Workers = append(mr.Workers, args.Worker)
	//广播给所有
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
			mr.NewCond.Wait()
		}
		mr.Unlock()
	}
}

//分布式执行mapReduce任务,通过rpc在主服务器上注册的workers调度map和reduce任务
func Distributed(jobName string, files []string, nReduce int, master string) *Master {
	fmt.Println("distributed....")
	mr := NewMaster(master)
	//启动master的RPC服务
	mr.StartRpcServer()
	//执行任务
	go mr.Run(jobName, files, nReduce, func(phase JobPhase) {
		ch := make(chan string)
		go mr.ForwardRegister(ch)
		//调度执行
		schedule(mr.JobName, mr.Files, mr.NReduce, phase, ch)
	}, func() {
		mr.Status = mr.KillWokre()
		mr.StopRpcServer()
	})
	return mr
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
