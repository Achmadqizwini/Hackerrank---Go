package main

import (
	"fmt"
	"strings"
)

func CountingValleys(steps int, path string) (valleyCounter int) {
	pathList := strings.Split(path, "")
	elev := 0
	var lembah bool

	for i := 0; i < steps; i++ {
		if pathList[i] == "U" {
			elev++
		} else {
			elev--
		}
		if elev < 0 && lembah == false {
			lembah = true
			valleyCounter++
		}
		if elev >= 0 && lembah == true {
			lembah = false
		}
	}
	return valleyCounter
}

func main() {
	a := "abcdefghijk"
	b := string(a[0])
	fmt.Println(b)
}
