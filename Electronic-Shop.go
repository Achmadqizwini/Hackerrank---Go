package main

import "fmt"

func getMoneySpent(keyboards []int32, drives []int32, b int32) int32 {
	/*
	 * Write your code here.
	 */
	var max int32
	for i := 0; i < len(keyboards); i++ {
		for j := 0; j < len(drives); j++ {
			var a = keyboards[i] + drives[j]

			if a > max && a <= b {
				max = a
			}

		}
		if max == 0 {
			max = -1
		}
	}
	return max
}

func main() {
	fmt.Println(getMoneySpent([]int32{4}, []int32{5}, 5))
	fmt.Println(getMoneySpent([]int32{3, 1}, []int32{5, 2, 8}, 10))
	fmt.Println(getMoneySpent1([]int32{40, 50, 60}, []int32{5, 8, 12}, 60))

}
