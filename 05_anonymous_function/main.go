package main

import "fmt"

func main() {
	i := 8
	bob := func() int {
		i ++
		return i
	}

	fmt.Println(bob())
	fmt.Println(bob())
}
