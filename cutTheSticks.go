package main

import (
	"fmt"
	"sort"
)

func cutTheSticks(arr []int32) []int32 {
	// Write your code here
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	var sticks_cut []int32
	var res []int32
	var numofstick int
	for len(arr) != 0 {
		for i := 0; i < len(arr); i++ {
			a := arr[i] - arr[0]
			if a == 0 {
				continue
			}
			res = append(res, a)
			numofstick = len(res)
			sticks_cut = append(sticks_cut, int32(numofstick))
		}
		arr = res
	}
	return sticks_cut
}

func main() {
	fmt.Println(cutTheSticks([]int32{1, 2, 3, 4, 3, 3, 2, 1}))
}
