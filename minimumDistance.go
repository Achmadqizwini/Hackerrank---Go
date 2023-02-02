package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func minimumDistances(a []int32) int32 {
	// Write your code here
	for(let i = 0; i<a.length; i++){
		if(a.indexOf(a[i]) !== a.lastIndexOf(a[i])){
		  let minDistance = a.lastIndexOf(a[i]) - a.indexOf(a[i]);
		  distances.push(minDistance);
		}
	}
	var disArr = []int32{}
for i, v := range a {
	
}
	if len(disArr) == 0 {
		return -1
	}
	var min = disArr[0]

	for i := 1; i < len(disArr); i++ {
		if min > disArr[i] {
			min = disArr[i]
		}
	}
	return min
}

func main() {
	// fmt.Println(minimumDistances([]int32{7, 1, 3, 4, 7, 5, 8}))

}
