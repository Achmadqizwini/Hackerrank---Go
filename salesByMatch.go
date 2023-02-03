package main

import (
	"fmt"
)

func sockMerchant(n int, ar []int32) int32 {
	// Write your code here
	temp := make(map[int32]int)
	for _, v := range ar {
		temp[int32(v)]++
	}

	var res int32
	for _, v := range temp {
		a := int32(v / 2)
		res += a
	}
	return res
}

func main() {
	fmt.Println(sockMerchant(9, []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}))
}
