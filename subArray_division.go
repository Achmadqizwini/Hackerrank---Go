package main

import "fmt"

func birthday(s []int32, d int32, m int32) int32 {
	// Write your code here
	var sum int32
	var res int32
	for i := 0; i < len(s)-(int(m)-1); i++ {

		for j := i; j < i+int(m); j++ {
			sum += s[j]
		}

		if sum == d {
			res++
		}

		sum = 0
	}
	return res
}

func main() {
	fmt.Println(birthday([]int32{2, 2, 1, 3, 2}, 4, 2))
	fmt.Println(birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
}
