package main

import "fmt"

func main() {
	var n1 int
	var n2 int
			var n3 int
				var n4 int

	fmt.Println("рамадины топ?")
	fmt.Println("Введите вариант ответа 1.да/2.нет")
	fmt.Scanf("%d\n",&n1)
fmt.Println("я крутой")
		fmt.Println("Введите вариант ответа 1.да/2.нет")
	fmt.Scanf("%d\n",&n2)
fmt.Println("сколько лет")
		fmt.Println("Введите вариант ответа 1.да/2.нет")
	fmt.Scanf("%d\n",&n3)
	fmt.Println("что значит столовать")
		fmt.Println("Введите вариант ответа 1.да/2.нет")
	fmt.Scanf("%d\n",&n4)

	test(n1,n2,n3,n4)
}

func test(n1 int,n2 int,n3 int,n4 int){
if n1 == 1{
	fmt.Println("1.молодец")
} else if n1 == 2{
	fmt.Println("1.кака")
	}else{ fmt.Println("error")
}

if n2 == 1{
	fmt.Println("2.молодец")
} else if n1 == 2{
	fmt.Println("2.кака")
}else{
fmt.Println("error")
} 

if n3 == 2{
	fmt.Println("3.молодец")
} else if n1 == 1{
	fmt.Println("3.кака")
}else{
fmt.Println("error")
} 

if n4 == 2{
	fmt.Println("4.молодец")
} else if n1 == 1{
	fmt.Println("4.кака")
}else{
fmt.Println("error")
} 
}