package main

import "fmt"

var x = 42 // x is available at the package level.

func main() {
	fmt.Printf("regular string")

	p1 := person{
		first: "James",
	}
	p1.sayHello()
}

type person struct {
	first string
}

func (p person) sayHello() {
	fmt.Println("hello my name is ", p.first)
}
