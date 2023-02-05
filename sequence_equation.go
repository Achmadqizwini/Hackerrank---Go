package main

import "fmt"

func permutationEquation(p []int32) []int32 {
	var temp1 = map[int32]int32{}
	var res = []int32{}

	for i := 1; i <= len(p); i++ {
		for j, v := range p {
			if v == int32(i) {
				temp1[int32(i)] = int32(j + 1)
			}
		}
	}
	for i := 1; i <= len(p); i++ {
		res = append(res, temp1[temp1[int32(i)]])
	}
	return res
}

func main() {
	fmt.Println(permutationEquation([]int32{5, 2, 1, 3, 4}))
}
