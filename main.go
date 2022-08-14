package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

		var outData [][]string

		matrixLen := len(dv)
		for i := 0; i < matrixLen; i++ {
			//d0 就是一串串的字符
			var d0 = dv[i]

			if i == 0 {
				for _, v := range d0 {
					outData = append(outData, []string{v})
				}
			}
			//构建第1列
		}
		fmt.Printf("%v\n", outData)
	} else {

		flag.Usage()
	}
}
