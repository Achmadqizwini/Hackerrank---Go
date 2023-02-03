package main

import "fmt"

func angryProfessor(k int32, ar []int32) string {
	res := ""
	var temp int32
	for _, v := range ar {
		if v <= 0 {
			temp++
		}
	}
	if temp >= k {
		res = "NO"
	} else {
		res = "YES"
	}
	return res
}

func main() {
	fmt.Println(angryProfessor(3, []int32{-1, -3, 4, 2}))
	fmt.Println(angryProfessor(2, []int32{0, -1, 2, 1}))
	fmt.Println(angryProfessor(4, []int32{0, -1, 2, 1}))
	fmt.Println(angryProfessor(4, []int32{0, -1, -2, -3}))
}
