package main

import (
	"fmt"
	"sort"
)

func hurdleRace(k int32, ar []int32) int32 {
	sort.SliceStable(ar, func(i, j int) bool {
		return ar[i] > ar[j]
	})
	var res = ar[0] - k
	if res < 0 {
		return 0
	}
	return res
}

func main() {
	fmt.Println(hurdleRace(1, []int32{1, 2, 3, 3, 2}))
	fmt.Println(hurdleRace(4, []int32{1, 6, 3, 5, 2}))
	fmt.Println(hurdleRace(7, []int32{2, 5, 4, 5, 2}))
}
