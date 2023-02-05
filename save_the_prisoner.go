package main

import "fmt"

func saveThePrisoner(n, m, s int32) int32 {
	var res int32
	a := (m + s - 1) % n
	if a == 0 {
		res = n
	} else {
		res = a
	}
	return res
}
func save2(n, m, s int32) int32 {
	a := (m % n) + s - 1
	return a
}

func main() {
	fmt.Println(saveThePrisoner(4, 6, 2))  //3
	fmt.Println(saveThePrisoner(4, 4, 2))  //1
	fmt.Println(save2(2, 739424390, 2))    // 4
	fmt.Println(saveThePrisoner(5, 2, 1))  //2
	fmt.Println(saveThePrisoner(5, 2, 2))  //3
	fmt.Println(saveThePrisoner(7, 19, 2)) //6
	fmt.Println(saveThePrisoner(3, 7, 3))  //3
	fmt.Println("============")
	fmt.Println(save2(4, 6, 2))         //3
	fmt.Println(save2(4, 4, 2))         //1
	fmt.Println(save2(2, 739424390, 2)) // 4
	fmt.Println(save2(5, 2, 1))         //2
	fmt.Println(save2(5, 2, 2))         //3
	fmt.Println(save2(7, 19, 2))        //6
	fmt.Println(save2(3, 7, 3))         //3

}
