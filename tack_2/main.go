package main

import "fmt"

func main() {
	var n1 int
	var znak string
	var n2 int
	fmt.Println("Введите первое число:")
	fmt.Scanf("%d\n", &n1)

	fmt.Println("Введите операцию :")
	fmt.Scanf("%s\n", &znak)

	fmt.Println("Введите второе число:")
	fmt.Scanf("%d\n", &n2)
	Calculate(znak, n1, n2)

}

func Calculate(znak string, n1 int, n2 int) {
	if znak == "+" {
		fmt.Println(n1 + n2)
	} else if znak == "-" {
		fmt.Println(n1 - n2)
	} else if znak == "/" {
		if n2 == 0 {
			fmt.Println("0")
		}else{
			fmt.Println(n1 / n2)
		}
	} else if znak == "*" {
		fmt.Println(n1 * n2)
	} else {
		fmt.Println("error")
	}
}
