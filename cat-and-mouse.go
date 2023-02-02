package main

import (
	"fmt"
	"math"
)

func catAndMouse(x, y, z int32) string {
	res := ""
	x1 := math.Abs(float64(z - y))
	y1 := math.Abs(float64(z - x))
	if x1 > y1 {
		res = "Cat A"
	} else if x1 < y1 {
		res = " Cat B"
	} else {
		res = "Mouse C"
	}
	return res
}

func main() {
	fmt.Println(catAndMouse(2, 5, 4))
	fmt.Println(catAndMouse(1, 2, 3))
	fmt.Println(catAndMouse(1, 3, 2))
}
