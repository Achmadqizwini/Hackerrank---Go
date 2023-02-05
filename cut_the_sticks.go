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
		fmt.Println(arr)
	}
	return sticks_cut
}

func cutSticks(arr []int32) []int32 {
	var newArr = []int32{}
	for {
		var min int32 = 1001

		length := 0
		n := len(arr)
		for i := 0; i < n; i++ {
			if arr[i] != 0 {

				length++
				if arr[i] < min {
					min = arr[i]
				}
			}

		}
		for i := 0; i < n; i++ {
			if arr[i] != 0 {
				arr[i] = arr[i] - min
			}

		}
		if length == 0 {
			break
		}
		fmt.Println(length)
		newArr = append(newArr, int32(length))
	}
	return newArr
}
func main() {
	// fmt.Println(cutTheSticks([]int32{1, 2, 3, 4, 3, 3, 2, 1}))
	// fmt.Println(cutTheSticks([]int32{1, 2, 3, 4, 3, 3, 2, 1}))
	fmt.Println(cutSticks([]int32{1, 2, 3, 4, 3, 3, 2, 1}))
}
