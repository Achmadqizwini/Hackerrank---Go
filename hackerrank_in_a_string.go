package main

import (
	"fmt"
	"strings"
)

func checkString(s2 string) string {
	s1 := "hackerrank"
	for _, c := range s1 {
		if strings.Count(s2, string(c)) < strings.Count(s1, string(c)) {
			return "NO"
		}
	}
	return "YES"
}
func checkSequence(s2 string) string {
	s1 := "hackerrank"
	i := 0
	for _, c := range s2 {
		if i < len(s1) && c == rune(s1[i]) {
			i++
		}
	}
	if i == len(s1) {
		return "YES"
	} else {
		return "NO"
	}
}
func main() {
	fmt.Println(checkString("hhaacckkekraraannk"))
	fmt.Println(checkString("hereiamstackerrank"))
	fmt.Println(checkString("hackerworld"))
	fmt.Println(checkString("rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt"))

	fmt.Println("============")
	fmt.Println(checkSequence("hhaacckkekraraannk"))
	fmt.Println(checkSequence("hereiamstackerrank"))
	fmt.Println(checkSequence("hackerworld"))
	fmt.Println(checkSequence("rhbaasdndfsdskgbfefdbrsdfhuyatrjtcrtyytktjjt"))
}
