package main

import (
	"errors"
	"fmt"
)

func main() {
	data, err := Divide(1, 0)
	if err != nil {
		fmt.Println("Errrrrrrrr")
		return
	}
	fmt.Println(data)
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("Err")
	}
	res := a / b
	return res, nil
}
