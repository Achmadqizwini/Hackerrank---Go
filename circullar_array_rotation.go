package main

import "fmt"

func circularArrayRotation(a []int32, k int32, queries []int32) []int32 {
	// Write your code here
	var res = []int32{}
	for i := 1; i <= int(k); i++ {
		var temp = []int32{}
		lastIndex := a[len(a)-1]
		b := a[:(len(a) - 1)]
		temp = append(temp, lastIndex)
		temp = append(temp, b...)
		a = temp
	}

	for _, v := range queries {
		res = append(res, a[v])
	}
	return res
}

func main() {
	fmt.Println(circularArrayRotation([]int32{3, 4, 5}, 2, []int32{1, 2}))
}
