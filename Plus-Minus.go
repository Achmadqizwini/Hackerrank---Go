package main

import "fmt"

func plusMinus(ar []int32) {
	var min, zero, plus int32
	for i := 0; i < len(ar); i++ {
		if ar[i] < 0 {
			min++
		} else if ar[i] == 0 {
			zero++
		} else {
			plus++
		}
	}
	minRatio := float64(min) / float64(len(ar))
	zeroRatio := float64(zero) / float64(len(ar))
	maxRatio := float64(plus) / float64(len(ar))

	fmt.Printf("%.6f\n %.6f\n %.6f\n", minRatio, zeroRatio, maxRatio)
}
func main() {
	plusMinus([]int32{-1, -2, 0, 1, 1})
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
}
