package main

import (
	"fmt"
	"math"
)

func diagonalDifference(ar [][]int32) int32 {
	var res, sumA, sumB int32
	for i, _ := range ar {
		sumA += ar[i][i]
		sumB += ar[i][len(ar)-1-i]
	}
	res = int32(math.Abs(float64((sumA) - (sumB))))
	return res
}
func main() {
	fmt.Println(diagonalDifference([][]int32{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}))
}
