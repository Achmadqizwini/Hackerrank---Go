package main

import "fmt"

func jumpingOnClouds(c []int32, k int32) int32 {
	var energy = 100
	for i := 0; i <= len(c); i = i + int(k) {
		if c[i] == 1 {
			energy = energy - 3
		} else {
			energy -= 1
		}
		fmt.Println("energyyyyy", i, energy)
		if i+int(k) > len(c) {
			i = (i + int(k)) - len(c) - int(k)
		}
		if (i+int(k))-len(c) == 0 {
			break
		}
		fmt.Println("new i", i)
	}
	return int32(energy)
}
func main() {
	// fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 1, 0}, 2))
	// fmt.Println(jumpingOnClouds([]int32{0, 0, 1, 0}, 2))
	fmt.Println(jumpingOnClouds([]int32{1, 1, 1, 0, 1, 1, 0, 0, 0, 0}, 3))

}
