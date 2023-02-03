package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	var res int32
	for _, v := range ar {
		res += v
	}
	return res
}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
}
