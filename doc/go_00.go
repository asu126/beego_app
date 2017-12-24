package main

import "fmt"

func addressStillTest(v *int) {
	x := 456
	v = &x
}

func main() {
	x := 1000
	fmt.Println("x ", x)
	addressStillTest(&x)
	fmt.Println("after x ", x)
}

/*
x  1000
after x  1000
*/
