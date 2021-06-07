package mapReduce

import "net/rpc"

type ShutDownReplay struct {
	//代表指定woker执行的当前为止的任务数量（任务编号）
	NTasks int
}

//实现rpc请求发送函数
/*
	svc: 地址
	rpcName：服务方法
	args： 传递参数
	reply： 参数
*/
func call(svc string, rpcName string, args interface{}, reply interface{}) bool {
	//链接rpc服务
	c, err := rpc.Dial("unix", svc)
	if err != nil {
		return false
	}
	defer func() {
		_ = c.Close()
	}()
	//调用指定方法
	err = c.Call(rpcName, args, reply)
	if err == nil {
		return true
	}
	return false
}

//任务参数结构，存储任务相关的信息
type DoTaskArgs struct {
	JobName   string
	File      string   //输入文件只对map有用
	Phase     JobPhase //任务类型
	TaskNumer int      // 任务编号
	//map需要改参数计算中间结果文件的输出数量
	//reduce需要该参数获取收集中间结果文件数量
	NumOtherPhase int //另一个任务类型的任务总数
}
