package main

import (
	"fmt"
	"sort"
)

func getTotalX(a []int32, b []int32) int32 {
	// Write your code here
	var factor = []int32{}
	var count int32
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	for i := a[len(a)-1]; i <= (b[0]); i++ {
		if b[0]%int32(i) == 0 {
			factor = append(factor, int32(i))
			fmt.Println(factor)
		}
	}

	var fact int32
	var res = []int32{}
	for _, v := range factor {
		for _, val := range b {
			if val%v == 0 {
				fact++
			} else if val%v != 0 {
				break
			}
		}
		if fact == int32(len(b)) {
			res = append(res, v)
		}
		fact = 0
		fmt.Println("res:", res)
	}
	var num int32
	var result = []int32{}
	for _, v := range res {
		fmt.Println(v)
		for _, val := range a {
			fmt.Println("value a:", val)
			if v%val == 0 {
				num++
			}
		}
		if num == int32(len(a)) {
			fmt.Println("nummmmmm", num)
			result = append(result, v)
		}
		num = 0
	}
	count = int32(len(result))
	return count
}

func main() {
	// fmt.Println(getTotalX([]int32{2, 4}, []int32{16, 32, 96}))
	fmt.Println(getTotalX([]int32{3, 4}, []int32{24, 48}))
	// fmt.Println(getTotalX([]int32{2}, []int32{20, 30, 12}))
}
