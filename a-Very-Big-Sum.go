package main

import "fmt"

func aVeryBigSum(a []int64) int64 {
	var res int64
	for _, v := range a {
		res += v
	}
	return res
}

func main() {
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}
