package main

import "fmt"

func timeConversion(s string) string {
	// Write your code here
	var hour, minute, second int
	var ampm string
	fmt.Sscanf(s, "%d:%d:%d%s", &hour, &minute, &second, &ampm)
	if ampm == "PM" && hour != 12 {
		hour += 12
	}
	if ampm == "AM" && hour == 12 {
		hour = 0
	}
	return fmt.Sprintf("%02d:%02d:%02d", hour, minute, second)
}

func main() {
	fmt.Println(timeConversion("01:00:00PM"))
	fmt.Println(timeConversion("12:10:00PM"))
	fmt.Println(timeConversion("12:10:00AM"))
}
