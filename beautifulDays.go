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
		fmt.Println(a)
		fmt.Println("==============")
		for _, v := range a {
			result = string(v) + result
		}
		a2, _ := strconv.Atoi(result)
		fmt.Println(a2)
		fmt.Println("================")
		if (i-int32(a2))%k == 0 {
			res++
			fmt.Println(res)
		}

		result = ""
	}
	return int32(res)
}

func main() {
	fmt.Println(beautifulDays(20, 23, 6))
}
