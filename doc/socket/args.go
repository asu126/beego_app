package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("os.Args[0]")
	fmt.Println(os.Args[0])
	fmt.Println("os.Args[1]")
	fmt.Println(os.Args[1])
	fmt.Println("os.Args[2]")
	fmt.Println(os.Args[2])
}

/*
 go run args.go 1 2 3
*/
