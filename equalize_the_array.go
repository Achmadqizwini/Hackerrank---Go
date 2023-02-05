package main

import (
	"fmt"
	"sort"
)

func equalizeTheArray(ar []int32) int32 {
	var temp = map[int32]int32{}
	for _, v := range ar {
		temp[v]++
	}

	type x struct{ num, freq int32 }
	var sortMap []x

	for i, v := range temp {
		sortMap = append(sortMap, x{num: i, freq: v})
	}

	sort.SliceStable(sortMap, func(i, j int) bool {
		return sortMap[i].freq > sortMap[j].freq
	})
	fmt.Println(sortMap)

	var res int32
	for i := 1; i < len(sortMap); i++ {
		res += sortMap[i].freq
	}
	return res
}
func main() {
	fmt.Println(equalizeTheArray([]int32{3, 3, 2, 1, 3}))
	fmt.Println(equalizeTheArray([]int32{1, 2, 2, 3}))
}
