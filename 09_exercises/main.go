package main

import "fmt"

func main() {
	var xyz uint64 = 1
  var i uint64 = 1
	for i = 1 ; i <= 100 ; i++ {
		if(i%3==0 || i%5==0) {
			xyz = xyz*i
			fmt.Printf("(%d,%d), ",i,xyz)
		}
	}
	fmt.Println("")
	fmt.Printf("result = %d" , xyz)
}
