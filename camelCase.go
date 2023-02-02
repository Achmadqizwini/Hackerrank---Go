package main

import "fmt"

func camelcase(s string) int32 {
	var total int32 = 1

	for i := 1; i < len(s); i++ {

		if s[i] >= 65 && s[i] <= 90 {
			total += 1
		}
	}
	return total
}

func main() {
	fmt.Println(camelcase("satuDuaTigaEmpat"))
}
