package main

import "fmt"

func main() {
  //declare and initialize
	a := 10
	fmt.Printf("%T --> %v \n",a,a)

	b := "test"
	fmt.Printf("%T --> %v \n",b,b)

	var c string = "old style but i don't like it!"
	fmt.Printf("%T --> %v \n",c,c)

  //default go value
	var d float64
	fmt.Printf("%T --> %v \n",d,d)
}
