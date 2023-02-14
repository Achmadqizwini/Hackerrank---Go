package main

import "fmt"

func dayOfProgrammer(year int32) string {
	month := 9
	days := 13
	if year > 1918 {
		if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
			days--
		}
	} else if year == 1918 {
		days = 26
	} else {
		if year%4 == 0 {
			days--
		}
	}
	a := fmt.Sprintf("%02d.%02d.%4d\n", days, month, year)
	return a
}
func main() {
	fmt.Println(dayOfProgrammer(2017))
}
