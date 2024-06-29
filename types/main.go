package main

import "fmt"

func main() {
	x := 2
	y := 2.0
	fmt.Printf("%v of type %T \n", x, x)
	fmt.Printf("%v of type %T \n", y, y)

	y = float64(x)
	fmt.Printf("%v of type %T \n", y, y)

	const r = 10
}
