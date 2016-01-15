package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2

	p2 = "constant2"
	p3 = "constant3"
)
const p = "constant"

func main() {
	fmt.Println("p - ", p)
	fmt.Println("p2 - ", p2)
	fmt.Println("p3 - ", p3)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
