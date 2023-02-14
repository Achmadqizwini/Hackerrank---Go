package main

import (
	"fmt"
	"strconv"
)

func fairRations(B []int32) string {
	var odd int
	var res int
	for _, v := range B {
		if v%2 != 0 {
			odd++
		}
	}
	if odd%2 != 0 {
		return "NO"
	} else {
		for i, v := range B {
			if v%2 != 0 {
				v++
				B[i+1]++
				res += 2
			}
		}
	}
	result := strconv.Itoa(res)
	return result
}

func main() {
	fmt.Println(fairRations([]int32{4, 5, 6, 7}))
	fmt.Println(fairRations([]int32{2, 3, 4, 5, 6}))
	fmt.Println(fairRations([]int32{1, 2}))
}
