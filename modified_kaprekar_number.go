package main

import (
	"fmt"
	"strconv"
	"strings"
)

func kaprekarNumber(p, q int32) {
	var exist bool
	for i := p; i < q; i++ {
		sqr := i * i
		fmt.Println("sqr", sqr)
		a := strconv.Itoa(int(sqr))

		arrA := strings.Split(a, "")
		mid := len(a) / 2
		x := arrA[0:mid]
		y := arrA[mid:]

		fmt.Println("x", x)
		fmt.Println("y", y)

		x1 := ""
		y1 := ""

		for _, v := range x {
			x1 += v
		}
		for _, v := range y {
			y1 += v
		}

		x2, _ := strconv.Atoi(x1)
		y2, _ := strconv.Atoi(y1)

		if sum := x2 + y2; sum == int(i) {
			fmt.Printf("%d ", i)
			exist = true
		}
	}
	if exist == false {
		fmt.Println("INVALID RANGE")
	}
}

func main() {
	kaprekarNumber(1, 100)
	kaprekarNumber(77777, 77780)
}
