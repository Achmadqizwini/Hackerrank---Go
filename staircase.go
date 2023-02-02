package main

import "fmt"

func staircase(n int32) {
	sharp := ""
	for i := int32(0); i < n; i++ {
		for j := n - 1; j > i; j-- {
			sharp += (" ")
		}
		for k := int32(0); k <= i; k++ {
			sharp += "# "
		}
		sharp += "\n"
	}
	fmt.Println(sharp)
}
func staircase2(n int32) {
	// Write your code here
	sharp := ""
	for i := int32(0); i < n; i++ {
		for j := n - 1; j > i; j-- {
			sharp += " "
		}

		for k := int32(0); k <= i; k++ {
			sharp += "#"
		}
		sharp += "\n"
	}
	fmt.Println(sharp)
}

func staircase3(n int) {
	sharp := ""
	for i := 0; i < n; i++ {
		for j := n - 1; j >= i; j-- {
			sharp += "# "
		}
		sharp += "\n"
	}
	fmt.Println(sharp)
}

func staircase4(n int) {
	sharp := ""
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sharp += "#"
		}
		sharp += "\n"
	}
	fmt.Println(sharp)
}
func staircase5(n int) {
	sharp := ""
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sharp += " "
		}
		for k := n; k > i; k-- {
			sharp += "#"
		}
		sharp += "\n"
	}
	fmt.Println(sharp)
}
func staircase6(n int) {
	sharp := ""
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			sharp += " "
		}
		for k := n; k > i; k-- {
			sharp += "# "
		}
		sharp += "\n"
	}
	fmt.Println(sharp)
}
func main() {
	staircase(5)
	staircase2(5)
	staircase3(5)
	staircase4(5)
	staircase5(5)
	staircase6(5)
}
