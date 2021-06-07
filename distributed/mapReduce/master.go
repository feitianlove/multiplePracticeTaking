package mapReduce

type Master struct {
}

/*
	任务调度函数
	files 要处理的文件
	nReduce 分区数量
*/
func SequentialOne(jobName string, files []string, nReduce int,
	mapF func(string, string) []KeyValue,
	reduceF func(string, []string) string) {

	m := NewMaster()
	m.Run(jobName, files, nReduce, func(phase JobPhase) {
		switch phase {
		case mapPhase:
			for i, f := range files {
				DoMpa(jobName, i, f, nReduce, mapF)
			}
		case ReducePhase:
			for i := 0; i < nReduce; i++ {
				DoReduce(jobName, i, mergeName(jobName, i), len(files), reduceF)
			}
		}
	})
}

func NewMaster() *Master {
	return &Master{}
}

func (m *Master) Run(jobName string, files []string, nReduce int, schedule func(phase JobPhase)) {
	//执行map任务
	schedule(mapPhase)
	//执行reduce热卖
	schedule(ReducePhase)
	//合并文件
	m.merge(nReduce, jobName)
}
