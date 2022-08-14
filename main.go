package main

import (
	"encoding/json"
	"flag"
	"os"
)

var fn = flag.String("f", "input.json", "输入input.json")

func main() {
	flag.Parse()
	var dv [][]string
	if *fn != "" {
		data, err := os.ReadFile(*fn)
		if err != nil {
			panic("文件读取失败")

		}

		err = json.Unmarshal(data, &dv)
		if err != nil {
			panic("解析文件失败")
		}

		var outData = make([][]string, 0)

		for i := 0; i < len(dv); i++ {
			//d0 就是一串串的字符
			var d0 = dv[i]
			for _, v := range d0 {
				if outData[i] == nil {
					tmpOd := make([]string, 0)
					outData[i] = tmpOd
				}
				outData[i] = append(outData[i], v)
			}
			//构建第1列
		}
	} else {

		flag.Usage()
	}
}
