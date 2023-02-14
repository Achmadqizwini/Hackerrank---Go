package main

import (
	"fmt"
	"math"
)

func squares(a int32, b int32) int32 {
	// Write your code here
	var res int32
	for i := math.Ceil(math.Sqrt(float64(a))); i <= math.Floor(math.Sqrt(float64(b))); i++ {
		fmt.Println("i", i)
		res++
	}
	return res
}

func main() {
	fmt.Println(squares(3, 9))
	fmt.Println(squares(17, 24))
}
