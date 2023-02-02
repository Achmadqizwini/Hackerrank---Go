package main

import "fmt"

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var apel, jeruk int32
	for _, v := range apples {
		if v+a >= s && v+a <= t {
			apel++
		}
	}
	for _, v := range oranges {
		if v+b <= t && v+b >= s {
			jeruk++
		}
	}
	fmt.Println(apel)
	fmt.Println(jeruk)

}
func main() {
	countApplesAndOranges(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6})
}
