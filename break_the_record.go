package main

import "fmt"

func breakingRecords(ar []int32) []int32 {
	var min, max = ar[0], ar[0]
	var sumA, sumB int32
	var res []int32
	for i := 1; i < len(ar); i++ {
		if ar[i] < min {
			min = ar[i]
			sumA++
		} else if ar[i] > max {
			max = ar[i]
			sumB++
		}
	}
	res = append(res, sumB, sumA)
	return res
}

func main() {
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
}
