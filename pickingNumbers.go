package main

import (
	"fmt"
	"math"
	"sort"
)

func pickingNumbers(a []int32) int32 {
	// Write your code here
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var result int32
	var max = int32(0)
	for i := 0; i < len(a)-1; i++ {
		op := math.Abs(float64(a[i] - a[i+1]))
		po := math.Abs(float64(a[i] - a[i-1]))
		fmt.Println(op)
		if op <= 1 && po <= 1 {
			result++
			// 1,1,2,2,2,3
			// fmt.Println(result)
		}
		if result > max {
			max = result
		}

	}
	return max
}

func main() {
	// fmt.Println(pickingNumbers([]int32{4, 6, 5, 3, 3, 1}))
	fmt.Println(pickingNumbers([]int32{1, 2, 2, 3, 1, 2}))
}
