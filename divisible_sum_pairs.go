package main

import "fmt"

func divisibleSumPairs(n, k int32, ar []int32) int32 {
	var res int32
	for i, v := range ar {
		for j := i + 1; j < len(ar); j++ {
			if (v+ar[j])%k == 0 {
				res++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(divisibleSumPairs(6, 3, []int32{1, 3, 2, 6, 1, 2}))
}
