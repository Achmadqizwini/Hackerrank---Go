package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	var res []int32
	var sumA int32
	var sumB int32
	for i, _ := range a {
		if a[i] > b[i] {
			sumA++
		} else if a[i] < b[i] {
			sumB++
		}
	}
	res = append(res, sumA, sumB)
	return res
}

func main() {
	fmt.Println(compareTriplets([]int32{1, 3, 5}, []int32{2, 3, 1}))
}
