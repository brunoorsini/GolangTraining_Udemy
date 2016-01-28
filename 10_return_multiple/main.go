package main

import "fmt"

func main() {
	fmt.Println(greet("Bob ", "Léponge "))
}

func greet(firstname, lastname string) (string, string) {
	return fmt.Sprint(firstname, lastname), fmt.Sprint(lastname, firstname)
}
