package main

import "fmt"

type group struct {
	id   int
	name string
	age  int
	ms   []ms
}
type ms struct {
	con string
}

func main() {
	var n [3]group = [3]group{
		group{name: "Savely", age: 123, ms: []ms{
			ms{con: "ads"},
		}},
		group{name: "Margorita", age: 12},
		group{name: "Krutoy", age: 21}}

	for i := 0; i < len(n); i++ {
		n[i].id = i
		fmt.Println("Юзер по имени: ", n[i].name)
		fmt.Println("Его возраст : ", n[i].age)
		for i := 0; i < len(n[i].ms); i++ {

		}
	}
}
