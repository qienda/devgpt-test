package main

import "fmt"

func main() {
    var x int = "Hello, World!" // 这里将字符串赋值给一个整数变量，会导致类型错误
    fmt.Println(x)
}
