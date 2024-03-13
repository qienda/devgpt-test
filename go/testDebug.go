package main

import "fmt"

func main() {
slice := []int{1, 2, 3}
if len(slice) > 2 {
    fmt.Println(slice[2])

} else {
    fmt.Println("Index out of range")
}
 
}
