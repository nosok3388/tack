package main

import "fmt"

func main() {
    var ptr *int
    *ptr = 42 // ошибка?
    fmt.Println(*ptr)
}