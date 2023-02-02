package main

import (
	"fmt"
	"sort"
)

func migratoryBirds(ar []int32) int32 {
	var temp = map[int32]int32{}
	for _, v := range ar {
		temp[v]++
	}
	var max int32
	for _, v := range temp {
		max = v
		break
	}
	for _, v := range temp {
		if v > max {
			max = v
		}
	}
	var maxNum = []int32{}
	for i, v := range temp {
		if v == max {
			maxNum = append(maxNum, i)
		}
	}
	fmt.Println("max", maxNum)
	sort.SliceStable(maxNum, func(i, j int) bool {
		return maxNum[i] < maxNum[j]
	})
	res := maxNum[0]
	return res
}

func main() {
	fmt.Println(migratoryBirds([]int32{1, 1, 2, 2, 3}))
	fmt.Println(migratoryBirds([]int32{1, 2, 3, 4, 5, 4, 3, 2, 1, 3, 4}))
}
