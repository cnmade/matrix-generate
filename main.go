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

		// 输出的新二维数组，每个Element是固定长度的，长度等于  dv len, 比如4列

		// 但是这个 总的 数量，等于 第一列size 乘以第二列size 乘 第三列size 乘以 第四列size

		totalSize := 0
		for i := 0; i < matrixLen; i++ {
			//d0 就是一串串的字符
			var d0 = dv[i]

			if i == 0 {
				totalSize = len(d0)
				for _, v := range d0 {
					outData = append(outData, []string{v})
				}
			} else {
				totalSize = totalSize * len(dv[i])
			}

			//构建第1列
		}
		fmt.Printf("%v\n, totalSize: %v", outData, totalSize)
		for j := 0; j < totalSize; j++ {
			//组合穷举数组
		}
	} else {

		flag.Usage()
	}
}
