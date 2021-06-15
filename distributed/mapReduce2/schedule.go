package mapReduce

import (
	"fmt"
	"sync"
)

//调度函数的实现、决定如何向worker分配任务

func schedule(jobName string, mapFiles []string, nReduce int, phase JobPhase, registerChan chan string) {
	fmt.Println(registerChan)
	//当前任务数量
	var nTasks int
	//另一个任务数量
	var nOther int
	switch phase {
	case mapPhase:
		nTasks = len(mapFiles)
		nOther = nReduce
	case ReducePhase:
		nTasks = nReduce
		nOther = len(mapFiles)

	}
	//执行mapReduce的函数功能调用操作
	//所有任务都需要调度给works,等待所有任务执行成功之后在下一步，一个worker可以执行多个任务
	fmt.Println("schedule  jobName", jobName, phase, nTasks)
	//生成一个任务列表，将所有待处理的任务添加进去
	var wg sync.WaitGroup
	var lock *sync.Mutex = &sync.Mutex{}
	tasks := make([]int, nTasks)
	for i := 0; i < nTasks; i++ {
		tasks[i] = i
	}
	fmt.Println("schedule  tasks", tasks)
	var count int
	for {
		fmt.Println("------------------")
		lock.Lock()
		//如果没有任务执行
		if len(tasks) <= 0 {
			lock.Unlock()
			fmt.Println("sdfjaksdfjklasdfjlasdjflka")
			break
		}
		//	执行任务，每执行一个任务都要将任务从任务列表中删除
		task := tasks[0]
		if len(tasks) != 1 {
			tasks = append(tasks[:0], tasks[1:]...)
		} else {
			tasks = []int{}
		}
		lock.Unlock()
		//任务参数附值
		var doTaskArgs *DoTaskArgs = &DoTaskArgs{
			JobName:       jobName,
			File:          "",
			Phase:         phase,
			TaskNumer:     task,
			NumOtherPhase: nOther,
		}
		if phase == mapPhase {
			doTaskArgs.File = mapFiles[task]
		}
		worker := <-registerChan
		wg.Add(1)
		count++
		//调用rpc

		go func() {
			ok := call(worker, "Worker.DoTask", doTaskArgs, nil)
			fmt.Printf("Worker.DoTask%+v\n", doTaskArgs)
			fmt.Println("Worker.DoTask return", ok, tasks)
			//TODO
			wg.Done()
			if ok {
				registerChan <- worker
			} else {
				//worker执行失败，将该任务重新加入列表
				lock.Lock()
				tasks = append(tasks, doTaskArgs.TaskNumer)
				lock.Unlock()
			}
			//wg.Done()
			count--
		}()
	}
	fmt.Println("00000", count)

	wg.Wait()
	fmt.Printf("schedule %v finish\n", phase)
}
