package main

import "fmt"

func gradingStudents(ar []int32) []int32 {
	var res []int32
	for _, v := range ar {
		if v < 38 || v%5 == 1 || v%5 == 2 || v%5 == 0 {
			res = append(res, v)
		} else if v%5 == 3 {
			v = v + 2
			res = append(res, v)
		} else if v%5 == 4 {
			v = v + 1
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(gradingStudents([]int32{73, 67, 38, 33}))
}
