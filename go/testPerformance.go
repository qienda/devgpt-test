package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	// 创建一个非常大的切片
	slice := make([]int, 1000000)

	// 对切片进行连续的追加操作
	for i := 0; i < 1000000; i++ {
		slice = append(slice, i)
	}

	// 遍历切片并执行一些操作
	for _, value := range slice {
		// 模拟一些耗时的操作
		time.Sleep(time.Millisecond)
		fmt.Println(value)
	}

	elapsed := time.Since(start)
	fmt.Printf("程序执行时间：%s\n", elapsed)
}
