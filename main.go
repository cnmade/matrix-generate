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
			} else {
				totalSize = totalSize * len(dv[i])
			}

			//构建第1列
		}
		fmt.Printf("%v\n, totalSize: %v", outData, totalSize)

		var idxArr = make([]int, len(dv))

		for j := len(outData); j < totalSize; {

			x := generateElement(dv, &idxArr)
			fmt.Printf("outData = %v, x = %v\n", outData, x)
			//	if !listContains(outData, x) {
			outData = append(outData, x)
			//	}
			j = len(outData)
			//组合穷举数组
		}
		fmt.Println("%v\n", outData)
	} else {

		flag.Usage()
	}
}

func listContains(data [][]string, x []string) bool {
	if len(data) == 0 {
		return false
	}
	for _, v := range data {
		if doCompare(v, x) {
			return true
		}
	}
	return false
}

func doCompare(v []string, x []string) bool {
	s, _ := json.Marshal(v)
	z, _ := json.Marshal(x)
	fmt.Printf(" s: %v, z: %v \n", string(s), string(z))
	if string(s) == string(z) {
		return true
	}
	return false
}

func generateElement(dv [][]string, idx *[]int) []string {

	var outString []string

	for i := 0; i < len(dv); i++ {
		value, nextInt := generateValue(dv[i], (*idx)[i])
		(*idx)[i] = nextInt
		outString = append(outString, value)
	}
	return outString
}

func generateValue(strings []string, idx int) (string, int) {
	nowIdx := idx
	var nextInt int
	fmt.Printf("%v, idx: %v\n", strings, idx)
	if nowIdx >= len(strings) {
		nextInt = 0
	} else {
		nextInt = nowIdx + 1
	}
	if nowIdx >= len(strings) {
		nowIdx -= 1
	}
	return strings[nowIdx], nextInt
}
