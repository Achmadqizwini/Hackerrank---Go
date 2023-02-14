package main

import (
	"fmt"
	"sort"
)

func acmTeam(topic []string) []int32 {
	// Write your code here
	var max int32
	var team int32
	var res = []int32{}
	var tempMax = []int32{}
	for i, v := range topic {
		for j := i + 1; j < len(topic); j++ {
			var temp int32
			for k, w := range v {
				if string(w) == "1" || string(topic[j][k]) == "1" {
					temp++
				}
			}
			if temp >= max {
				max = temp
				tempMax = append(tempMax, max)
			}
		}
	}
	sort.SliceStable(tempMax, func(i, j int) bool {
		return tempMax[i] > tempMax[j]
	})
	var maxNum = tempMax[0]
	for _, v := range tempMax {
		if v == maxNum {
			team++
		}
	}
	res = append(res, max, team)
	return res
}

func main() {
	// fmt.Println(acmTeam([]string{"10101", "11110", "00010"}))
	fmt.Println(acmTeam([]string{"10101", "11100", "11010", "00101"}))
}
