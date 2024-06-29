package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	var g int // 0 value is stored if not declared
	fmt.Println(g)

	var q bool // 0 value is stored if not declared
	fmt.Println(q)
}
