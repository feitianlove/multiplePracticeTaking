package mapReduce

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

/*
	多个reduce节点生成的最终结果文件输出到一个文件中进行汇总
	nReduce 节点数量
	jobName 节点名称
*/
func (mr *Master) merge(nReduce int, jobName string) {
	var result map[string]string = make(map[string]string)

	for i := 0; i < nReduce; i++ {
		fileName := mergeName(jobName, i)
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		decoder := json.NewDecoder(file)
		var kv KeyValue
		for decoder.More() {
			err := decoder.Decode(&kv)
			if err != nil {
				panic(err)
			}
			result[kv.Key] = kv.Value
		}
		file.Close()
	}

	var keys []string
	for key, _ := range result {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	//	新建输出文件
	out_file, err := os.Create("mrtmp" + jobName)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(out_file)
	for _, key := range keys {
		fmt.Fprintf(w, "%s:%s\n", key, result[key])
	}
	w.Flush()
	out_file.Close()
}
