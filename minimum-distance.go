package main

import (
	"fmt"
	"math"
)

func minDistance(a []int32) int32 {
	var temp = map[int32]int32{}
	for _, v := range a {
		temp[v]++
	}
	var pairNum = []int32{}
	for i, v := range temp {
		if v == 2 {
			pairNum = append(pairNum, i)
		}
	}
	var b = []int{}
	var c = []int32{}
	for _, v := range pairNum {
		for j, w := range a {
			if v == w {
				b = append(b, j)
			}
		}
		x := math.Abs(float64(b[0] - b[1]))
		c = append(c, int32(x))
		b = []int{}
	}

	if len(c) == 0 {
		return -1
	}

	min := c[0]
	for _, v := range c {
		if v < min {
			min = v
		}
	}
	return min
}

func minimumDistance(ar []int32) int32 {
	var distance = []int32{}
	for i, v := range ar {
		for j := i + 1; j < len(ar); j++ {
			if v == ar[j] {
				a := math.Abs(float64(j - i))
				distance = append(distance, int32(a))
			}
		}
		if len(distance) == 0 {
			return -1
		}
	}
	min := distance[0]
	for _, v := range distance {
		if v < min {
			min = v
		}
	}
	return min

	// masih salah di pair-nya
}

func main() {
	fmt.Println(minimumDistance([]int32{3, 2, 1, 2, 3}))
	// fmt.Println(minimumDistance([]int32{3, 2, 1, 2, 3, 3, 3}))
	fmt.Println(minimumDistance([]int32{7, 1, 3, 4, 1, 7}))
	fmt.Println(minimumDistance([]int32{1, 2, 3, 4, 10}))
	fmt.Println("===============================")
	fmt.Println(minDistance([]int32{3, 2, 1, 2, 3, 3, 3}))
	fmt.Println(minDistance([]int32{7, 1, 3, 4, 1, 7}))
	fmt.Println(minDistance([]int32{1, 2, 3, 4, 10}))

}
