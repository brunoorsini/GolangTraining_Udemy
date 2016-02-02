package main

import "fmt"

//pb to solve : https://projecteuler.net/problem=31


func main() {
  //final_sum := 200
  coins := [8]int{200, 100, 50, 20, 10 , 5, 2, 1}


  for index := 0; index < len(coins); index++ {
    fmt.Println(coins[index])
  }



}
