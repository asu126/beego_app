package main

import "fmt"

type Dog struct {
	Name string
	Type string
}

func addressTest(d *Dog) {
	a := &Dog{"another cute dog", "another type"}
	d = a
}

func anotherTest(d *Dog) {
	a := &Dog{"another cute dog", "another type"}
	d.Name = a.Name
	d.Type = a.Type
}

func anotherAddressTest(d **Dog) {
	a := &Dog{"address dog 8888", "address dog type6666"}
	*d = a
}

func main() {
	d1 := &Dog{"1", "2"}
	fmt.Printf("before-1:%v\r\n", d1)
	fmt.Printf("before-1:%v\r\n", *d1)
	addressTest(d1)
	fmt.Printf("after-1:%v\r\n", d1)
	fmt.Printf("after-1:%v\r\n", *d1)
	anotherTest(d1)
	fmt.Printf("after-2:%v\r\n", d1)
	fmt.Printf("after-2:%v\r\n", *d1)

	anotherAddressTest(&d1)
	fmt.Printf("after-3:%v\r\n", d1)
	fmt.Printf("after-3:%v\r\n", *d1)
}

/*
before-1:&{1 2}
before-1:{1 2}
after-1:&{1 2}
after-1:{1 2}
after-2:&{another cute dog another type}
after-2:{another cute dog another type}
*/
