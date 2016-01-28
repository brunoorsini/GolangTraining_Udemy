package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	fmt.Println("average of ",data," is ",average(data...))
}

func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
