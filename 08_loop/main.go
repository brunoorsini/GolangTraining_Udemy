package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i, ", ")
	}
	fmt.Println("\n---- ---- ----")

	j := 0
	for  {
		if j>=10 {
			break
		}
		fmt.Print(j, ", ")
		j++;
	}
	fmt.Println("\n---- ---- ----")
}
