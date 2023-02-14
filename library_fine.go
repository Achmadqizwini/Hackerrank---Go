package main

import "fmt"

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	var res int32
	if d1 > d2 && m1 == m2 && y1 == y2 {
		res = 15 * (d1 - d2)
	} else if m1 > m2 && y1 == y2 {
		res = 500 * (m1 - m2)
	} else if y1 > y2 {
		res = 10000 * (y1 - y2)
	}
	return res
}

func main() {
	fmt.Println(libraryFine(9, 6, 2015, 6, 6, 2015))
}
