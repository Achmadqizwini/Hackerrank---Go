package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(ar []int32) {
	sort.SliceStable(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})
	var min, max int32
	for i := 0; i < len(ar)-1; i++ {
		min += ar[i]
	}
	for i := 1; i < len(ar); i++ {
		max += ar[i]
	}
	fmt.Println(min, max)
}
func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}
