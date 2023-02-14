package main

import "fmt"

func workbook(n int32, k int32, arr []int32) int32 {
	// Write your code here
	var container = [][]int32{}
	var temp = []int32{}
	var res int32
	for _, v := range arr {
		for i := 1; i <= int(v); i++ {
			temp = append(temp, int32(i))
			if len(temp) == int(k) || i == int(v) {
				container = append(container, temp)
				temp = []int32{}
			}
		}
	}
	// for _, v := range container {
	// 	fmt.Println(v)
	// }

	// fmt.Println(container)

	for i := 0; i < len(container); i++ {
		for _, w := range container[i] {
			if i+1 == int(w) {
				res++

			}
			fmt.Println("i :", 1, "w :", w)
		}
	}

	return res
}
func main() {
	// fmt.Println(workbook(5, 3, []int32{4, 2, 6, 1, 10}))
	// fmt.Println(workbook(15, 20, []int32{1, 8, 19, 15, 2, 29, 3, 2, 25, 2, 19, 26, 17, 33, 22}))
	fmt.Println(workbook(1, 1, []int32{100}))
}
