package main

import "fmt"

func main() {
slice := []int{1, 2, 3}
if len(slice) > 3 {
    fmt.Println(slice[3])
} else {
    fmt.Println("Index out of range")
}
 // 这里访问了超出切片范围的索引，会导致 panic
}
