package main

import "fmt"

func utopianTree(n int32) int32 {
	// Write your code here
	var tree int32
	for i := 0; i <= int(n); i++ {
		if i == 0 {
			tree = 1
		} else if i%2 == 1 {
			tree = tree * 2
		} else if i%2 == 0 {
			tree += 1
		}
		fmt.Println(tree)
	}
	return tree
}

func main() {
	fmt.Println(utopianTree(3))
	fmt.Println(utopianTree(5))
	fmt.Println(5 / 2)
}
