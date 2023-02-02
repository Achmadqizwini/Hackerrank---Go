package main

import "fmt"

func designerPdfViewer(h []int32, word string) int32 {
	// Write your code here
	chara := "abcdefghijklmnopqrstuvwxyz"
	var result int32
	temp := map[string]int32{}
	var words = []int32{}

	for i := 0; i < len(chara); i++ {
		temp[string(chara[i])] = h[i]
	}
	fmt.Println(temp)

	for i, _ := range word {

		words = append(words, temp[string(word[i])])
	}
	fmt.Println(words)

	max := words[0]
	for i := 1; i < len(words); i++ {
		if words[i] > max {
			max = words[i]
		}
	}

	result = max * int32(len(words))

	fmt.Println(max)

	return result
}

func main() {
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "abc"))
	fmt.Println(designerPdfViewer([]int32{1, 3, 1, 3, 1, 4, 1, 3, 2, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 7}, "zaba"))
}
