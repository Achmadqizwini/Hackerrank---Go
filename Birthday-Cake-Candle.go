package main

import (
	"fmt"
	"sort"
)

func birthdayCakeCandles(ar []int32) int32 {
	var res int32
	var temp = map[int32]int32{}
	sort.SliceStable(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	for _, v := range ar {
		temp[v]++
	}
	max := ar[len(ar)-1]
	res = temp[max]
	return res
}

func birthdayCakeCandles1(candles []int32) int32 {
	// Write your code here
	sort.SliceStable(candles, func(i, j int) bool {
		return candles[i] < candles[j]
	})

	var result []int32
	var highestNumber int32 = 0

	for i := len(candles) - 1; i >= 0; i-- {
		if candles[i] >= highestNumber {
			highestNumber = candles[i]
			result = append(result, candles[i])
		}
	}

	return int32(len(result))
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3, 2, 2}))
	fmt.Println("==========================")
	fmt.Println(birthdayCakeCandles1([]int32{3, 2, 1, 3}))
}
