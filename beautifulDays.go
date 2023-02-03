package main

import (
	"fmt"
	"strconv"
)

func beautifulDays(i int32, j int32, k int32) int32 {
	// Write your code here
	result := ""
	var res int
	for i := i; i <= j; i++ {
		a := strconv.Itoa(int(i))

		for _, v := range a {
			result = string(v) + result
		}
		a2, _ := strconv.Atoi(result)

		if (i-int32(a2))%k == 0 {
			res++
		}

		result = ""
	}
	return int32(res)
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
