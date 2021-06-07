package mapReduce

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
	创建一个包含N个编号的输入文件
	通过mapReduce进行处理
	检查输出文件中是否包含了N个编号
*/
func main() {
	//SequentialOne("test", makeInputs(1), 1, MapFunc, ReduceFunc)
	SequentialOne("test", makeInputs(5), 3, MapFunc, ReduceFunc)
}

//自定义map func
func MapFunc(file string, value string) []KeyValue {
	var res []KeyValue
	words := strings.Fields(value)
	for _, work := range words {
		kv := KeyValue{
			Key:   work,
			Value: "",
		}
		res = append(res, kv)
	}
	return res
}

// 自定义reduce func
func ReduceFunc(key string, value []string) string {
	for _, element := range value {
		fmt.Printf("Reduce %s-%v\n", key, element)
	}
	return "jjj"
}

const (
	nNum = 100
)

// 创建输入文件，返回创建好的文件列表,写入相应的数据
func makeInputs(num int) []string {
	var input []string
	var f = 0
	for i := 0; i < num; i++ {
		input = append(input, fmt.Sprintf("824-mrinput-%d.txt", i))
		file, err := os.Create(input[i])
		if err != nil {
			fmt.Printf("create input file [%s] err:%s\n", input[i], err)
		}
		buf := bufio.NewWriter(file)
		for f < (i+1)*(nNum/num) {
			_, _ = fmt.Fprintf(buf, "%d\n", f)
			f++
		}
		_ = buf.Flush()
		_ = file.Close()
	}
	return input
}
