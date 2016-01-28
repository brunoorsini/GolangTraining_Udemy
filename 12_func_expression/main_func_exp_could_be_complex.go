package main

import "fmt"

func makeGreeter() func() string {
  a := 0
  return func() string {
    a += 1
		return fmt.Sprint("Hi " , a)
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
  fmt.Println(greet())

  greet2 := makeGreeter()
	fmt.Println(greet2())
  fmt.Println(greet2())
	fmt.Println(greet2())
  fmt.Println(greet2())
}
