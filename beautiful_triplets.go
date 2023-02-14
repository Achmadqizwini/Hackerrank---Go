package main

import "fmt"

func beautifulTriplets(d int32, arr []int32) int32 {
	var temp = map[int32]bool{}
	for _, v := range arr {
		temp[v] = true
	}
	var res int32
	for _, v := range arr {
		a := v + d
		if temp[a] == true {
			b := a + d
			if temp[b] == true {
				res++
			}
		}
	}
	return res
}
func main() {
	fmt.Println(beautifulTriplets(3, []int32{1, 2, 4, 5, 7, 8, 10}))
	fmt.Println(beautifulTriplets(1, []int32{2, 2, 3, 4, 5}))
}
