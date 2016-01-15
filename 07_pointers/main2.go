package main

import "fmt"

func zero(z *int) {
	fmt.Println("param in func z =", z)
  fmt.Println("param in func *z =", *z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println("x =",x)
  fmt.Println("&x =",&x)
	zero(&x)
	fmt.Println("x final =",x) // x is 0
}
