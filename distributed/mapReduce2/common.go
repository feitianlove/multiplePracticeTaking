package mapReduce

import (
	"hash/fnv"
	"strconv"
)

//任务类型
type JobPhase string

const (
	mapPhase    JobPhase = "Map"
	ReducePhase JobPhase = "Reduce"
)

//用于保存需要传递给map/reduce的key/value键值对

type KeyValue struct {
	Key   string
	Value string
}

//输出文件
func mergeName(jobName string, reduceTask int) string {
	return "mrtmp." + jobName + "-res-" + strconv.Itoa(reduceTask)
}

//中间文件名称
func ReduceName(jobName string, mapTaskNum int, reduceTask int) string {
	return "mrtmp." + jobName + "-" + strconv.Itoa(mapTaskNum) + "-" + strconv.Itoa(reduceTask)
}

// hash 函数
func iHash(s string) int {
	f := fnv.New32()
	_, _ = f.Write([]byte(s))
	return int(f.Sum32() & 0x7fffff)
}
