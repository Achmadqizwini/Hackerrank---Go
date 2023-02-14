package main

import "fmt"

func jumpOnClouds(c []int32) int32 {
	var sum int32
	for i := 0; i < len(c)-1; i++ {
		if i < len(c)-2 {
			if c[i+2] == 1 {
				sum++
			} else {
				sum++
				i += 1
			}
		} else {
			sum++
		}

	}
	return sum
}

func main() {
	fmt.Println(jumpOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
	fmt.Println(jumpOnClouds([]int32{0, 0, 0, 1, 0, 0}))
}
