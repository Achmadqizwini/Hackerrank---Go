package main

import (
	"fmt"
	"strings"
)

func repeatedString(s string, n int64) int64 {
	a := strings.Split(s, "")
	var num int64
	var arr = a
	for len(arr) < int(n) {
		if len(arr) == 1 {
			if arr[0] == "a" {
				return n
			} else {
				return 0
			}
		}
		for _, v := range a {
			arr = append(arr, v)
			if len(arr) == int(n) {
				break
			}
		}
	}
	for _, v := range arr {
		if v == "a" {
			num++
		}
	}
	fmt.Println(arr)
	return num
}

func repeatedString2(s string, n int64) int64 {
	var num int64

	if len(s) == 1 {
		if string(s[0]) == "a" {
			return n
		} else {
			return 0
		}
	}

	a := n / int64(len(s))
	b := n % int64(len(s))
	newString := strings.Repeat(s, int(a))
	sliceStr := strings.Split(newString, "")

	for i := 0; i < int(b); i++ {
		sliceStr = append(sliceStr, string(s[i]))
	}
	fmt.Println(sliceStr)

	for _, v := range sliceStr {
		if v == "a" {
			num++
		}
	}
	return num

}
func repeatedString3(s string, n int64) int64 {
	// Write your code here
	p := int64(len(s))
	count := int64(n / p)
	sisa := int64(n % p)
	result := int64(0)
	do := false
	for _, v := range s {
		if string(v) == "a" {
			do = true
			break
		}
	}

	if do {
		if len(s) == 1 {
			return n
		}
		sum := 0 // jumlah a pada s
		for _, v := range s {
			if string(v) == "a" {
				sum += 1
			}
		}

		result = count * int64(sum)
		if sisa > 0 {
			slisisa := s[:sisa]
			for _, v := range slisisa {
				if string(v) == "a" {
					result += 1
				}
			}
		}
		return result
	} else {
		return 0
	}
}

func repeatedString4(s string, n int64) int64 {
	if s == "a" {
		return n
	}

	var div = n / int64(len(s))
	var modulo = int(n) % (len(s))
	var res int64
	for key, v := range s {
		if string(v) == "a" && key < modulo {
			res += div + 1
			fmt.Println("div", div, "key:", key, "res: ", res)
		} else if string(v) == "a" {
			res += div
		}
	}

	return res
}
func main() {
	// fmt.Println(repeatedString("abaab", 7))
	// fmt.Println(repeatedString("x", 7))
	// fmt.Println(repeatedString("a", 1000000000000))
	// fmt.Println("===========================")
	// fmt.Println(repeatedString2("abaab", 7))
	// fmt.Println(repeatedString2("x", 7))
	// fmt.Println(repeatedString2("a", 1000000000000))
	fmt.Println("===========================")
	fmt.Println(repeatedString3("abaab", 7))
	fmt.Println(repeatedString3("aba", 10))
	fmt.Println(repeatedString3("x", 7))
	fmt.Println(repeatedString3("a", 1000000000000))
	fmt.Println("===========================")
	fmt.Println(repeatedString4("abaab", 7))
	fmt.Println(repeatedString4("aba", 10))
	fmt.Println(repeatedString4("x", 7))
	fmt.Println(repeatedString4("a", 1000000000000))
}
