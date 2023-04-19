package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./test.py")
	if err != nil {
		panic(err)
	}
	chunk := make([]byte, 1024) // 创建一个 1KB 大小的字节数组
	total := 0
	for {
		n, err := file.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Println(n)
		total = total + n
	}
	fmt.Println(string(chunk[:total]))
	output, err := os.Create("./output.py")
	if err != nil {
		panic(err)
	}
	defer output.Close() // 在函数结束时关闭文件句柄
	if _, err := output.Write(chunk[:total]); err != nil {
		panic(err)
	}
}
