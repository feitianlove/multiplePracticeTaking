package mapReduce

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	实现一个map任务管理函数，从input files中读取内容
	将输出分成指定数量的中间文件
	自定义分割标注
*/

// nReduce 当前map执行的reduce编号
/*
	从输入文件infile中读取内容
	通过mapF对内容进行处理，分割map输出
	将指定文件结果解析为key/value
	创建nReduce个中间文件名称
*/
func DoMpa(jobName string, mapTaskNum int, inFile string, nReduce int, mapF func(string, string) []KeyValue) {
	f, err := os.Open(inFile)
	if err != nil {
		fmt.Printf("raad file %s err: %v\n", inFile, err)

	}
	context, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("read the context fail%v\n", err)
	}
	res := mapF(inFile, string(context))
	encodes := make([]*json.Encoder, nReduce)
	for i := 0; i < nReduce; i++ {
		fileName := ReduceName(jobName, mapTaskNum, i)
		f, err := os.Create(fileName)
		if err != nil {
			fmt.Printf("create file [%s] error :%v\n", fileName, err)
		}
		defer f.Close()
		encodes[i] = json.NewEncoder(f)
	}
	for _, v := range res {
		index := iHash(v.Key) % nReduce
		if err := encodes[index].Encode(&v); err != nil {
			fmt.Printf("unable to write file %v\n", err)
		}
	}
}
