package main

import (
	"fmt"
	"strings"
)

func CountingValleys(steps int, path string) (valleyCounter int) {
	pathList := strings.Split(path, "")
	fmt.Println(pathList)
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
	fmt.Println(CountingValleys(8, "UDDDUDUU"))
}
