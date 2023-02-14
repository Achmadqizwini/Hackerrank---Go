package main

import "fmt"

func chocolateFeast(n int32, c int32, m int32) int32 {
	var wrapper int32
	var total int32
	wrapper = n / c
	total = wrapper

	for {
		if wrapper < m {
			break
		} else if wrapper >= m {
			temp := wrapper / m
			temp2 := wrapper % m
			total += temp
			wrapper = temp + temp2
		}
	}
	return total
}

func main() {
	fmt.Println(chocolateFeast(10, 2, 5))
	fmt.Println(chocolateFeast(12, 2, 4))
}
