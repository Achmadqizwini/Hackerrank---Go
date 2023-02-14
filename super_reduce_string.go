package main

import "fmt"

func superReducedString(s string) string {
	// Write your code here
	runes := []rune(s)
	fmt.Println(runes)

	for {
		// find the index of the first pair of adjacent letters
		pairIndex := -1
		for i := 0; i < len(runes)-1; i++ {
			if runes[i] == runes[i+1] {
				pairIndex = i
				break
			}
		}

		// if there are no more pairs, return the final string
		if pairIndex == -1 {
			if len(runes) == 0 {
				return "Empty String"
			}
			return string(runes)
		}

		// remove the pair of adjacent letters
		runes = append(runes[:pairIndex], runes[pairIndex+2:]...)
		fmt.Println("=======")
		fmt.Println(runes)
	}
}

func main() {
	fmt.Println(superReducedString("aaabccddd"))
	fmt.Println(superReducedString("abba"))
}
