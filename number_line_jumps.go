package main

import "fmt"

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	// Write your code here
	max := 10000
	var meet = "NO"
	for i := 0; i <= max; i++ {
		if (x1 + (v1 * int32(i))) == (x2 + (v2 * int32(i))) {
			meet = "YES"
		}
	}
	return meet
}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2))
}
