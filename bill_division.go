package main

import (
	"fmt"
)

func billDivision(ar []int32, k, b int32) {
	var actual int32
	for i, v := range ar {
		if i == int(k) {
			continue
		}
		actual += v
	}
	// fmt.Println("actual", actual)
	res := b - (actual / 2)
	if res == 0 {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(res)
	}
}

func main() {
	billDivision([]int32{3, 10, 2, 9}, 1, 12)
	billDivision([]int32{3, 10, 2, 9}, 1, 7)
}
