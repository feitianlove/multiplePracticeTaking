package mapReduce

import (
	"encoding/json"
	"fmt"
	"os"
)

/*
	管理reduce任务
	打开每一个中间文件
	提取内容
	把具有相同内容的key合并
	生成最终结果文件
*/

func DoReduce(jobName string, reduceTaskNum int, outFile string, nMap int, reduceF func(string, []string) string) {
	var result map[string][]string = make(map[string][]string)
	for i := 0; i < nMap; i++ {
		interFile := ReduceName(jobName, i, reduceTaskNum)
		f, err := os.Open(interFile)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		decoder := json.NewDecoder(f)
		var kv KeyValue
		for decoder.More() {
			err := decoder.Decode(&kv)
			if err != nil {
				panic(err)
			}
			//fmt.Printf("%+v\n", kv)
			result[kv.Key] = append(result[kv.Key], kv.Value)
		}
	}
	fmt.Printf("%+v\n", result)

	//把内容做相应的处理
	var keys []string
	for key, _ := range result {
		keys = append(keys, key)
	}
	//	新建输出文件
	out_file, err := os.Create(outFile)
	if err != nil {
		panic(err)
	}
	defer out_file.Close()
	encoder := json.NewEncoder(out_file)
	for _, key := range keys {
		encoder.Encode(KeyValue{key, reduceF(key, result[key])})
	}
}
