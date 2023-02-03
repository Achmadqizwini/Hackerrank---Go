package main

import (
	"fmt"
	"strconv"
)

func findDigits(n int32) int32 {
	var res int32 = 0
	a := strconv.Itoa(int(n))
	for _, v := range a {
		b, _ := strconv.Atoi(string(v))
		if b == 0 {
			continue
		} else if n%int32(b) == 0 {
			res++
		}
	}
	return res
}
func main() {
	fmt.Println(findDigits(12))
	fmt.Println(findDigits(1012))
	fmt.Println(findDigits(124))
	fmt.Println(findDigits(111))
}
