package main

import (
	"fmt"
	"sort"
)

func pickingNumber(ar []int32) int32 { //17,5. gagal 2 test case. dont know why
	var temp = map[int32]int32{}
	var res int32
	sort.SliceStable(ar, func(i, j int) bool {
		return ar[i] < ar[j]
	})

	for _, v := range ar {
		temp[v]++
	}
	fmt.Println(temp)
	var max = temp[0]
	for _, v := range temp {
		if v > max {
			max = v
		}
	}
	fmt.Println("max", max)
	for i, v := range temp {
		if v == max {
			if v+temp[i+1] >= v+temp[i-1] {
				res = v + temp[i+1]
			} else if v+temp[i+1] <= v+temp[i-1] {
				res = v + temp[i-1]
			} else if temp[i+1] == 0 || temp[i-1] == 0 {
				res = v
			}
		}
	}
	return res
}

func pickingNumbers2(a []int32) int32 {
	// Write your code here
	var temp = map[int32]int32{}
	for _, v := range a {
		temp[v]++
	}

	fmt.Println("temp", temp)
	type tuple struct{ freq, num int32 }
	var tuples []tuple
	for num, freq := range temp {
		tuples = append(tuples, tuple{num: num, freq: freq})
	}
	fmt.Println("tuples", tuples)

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].num < tuples[j].num
	})
	fmt.Println("sorted tuples", tuples)
	max := tuples[0].freq
	for i := 1; i < len(tuples); i++ {
		if freq := tuples[i].freq; freq > max {
			max = freq
		}
		fmt.Println("====", tuples[i-1].num)
		if diff := tuples[i-1].num - tuples[i].num; diff >= -1 && diff <= 1 {
			sum := tuples[i-1].freq + tuples[i].freq
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func main() {
	// fmt.Println(pickingNumber([]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}))
	// fmt.Println(pickingNumber([]int32{5, 5, 5, 5, 5, 5, 5, 5, 5}))
	// fmt.Println(pickingNumber([]int32{98, 3, 99, 1, 97, 2}))
	// fmt.Println(pickingNumber([]int32{4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2, 98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3, 4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2, 98, 4, 2, 1, 1, 9, 2, 4}))
	// fmt.Println(pickingNumber([]int32{9, 6, 13, 16, 5, 18, 4, 10, 3, 19, 4, 5, 8, 1, 13, 10, 20, 17, 15, 10, 6, 10, 13, 20, 18, 17, 7, 10, 6, 5, 16, 18, 13, 20, 19, 7, 16, 13, 20, 17, 4, 17, 8, 19, 12, 7, 17, 1, 18, 3, 16, 4, 5, 3, 15, 17, 6, 17, 14, 11, 11, 7, 11, 6, 15, 15, 12, 6, 17, 19, 8, 6, 13, 9, 10, 19, 14, 18, 7, 9, 11, 16, 11, 20, 4, 20, 10, 7, 8, 4, 2, 12, 11, 8, 12, 13, 19, 8, 8, 5}))
	// fmt.Println(pickingNumber([]int32{7, 12, 13, 19, 17, 7, 3, 18, 9, 18, 13, 12, 3, 13, 7, 9, 18, 9, 18, 9, 13, 18, 13, 13, 18, 18, 17, 17, 13, 3, 12, 13, 19, 17, 19, 12, 18, 13, 7, 3, 3, 12, 7, 13, 7, 3, 17, 9, 13, 13, 13, 12, 18, 18, 9, 7, 19, 17, 13, 18, 19, 9, 18, 18, 18, 19, 17, 7, 12, 3, 13, 19, 12, 3, 9, 13, 19, 12, 18, 13, 18, 18, 18, 17, 13, 3, 18, 19, 7, 12, 9, 18, 3, 13, 13, 9, 7}))

	fmt.Println("=============================")
	fmt.Println(pickingNumbers2([]int32{1, 1, 2, 2, 4, 4, 5, 5, 5}))
	fmt.Println(pickingNumbers2([]int32{5, 5, 5, 5, 5, 5, 5, 5, 5}))
	fmt.Println(pickingNumbers2([]int32{98, 3, 99, 1, 97, 2}))
	// fmt.Println(pickingNumbers2([]int32{4, 2, 3, 4, 4, 9, 98, 98, 3, 3, 3, 4, 2, 98, 1, 98, 98, 1, 1, 4, 98, 2, 98, 3, 9, 9, 3, 1, 4, 1, 98, 9, 9, 2, 9, 4, 2, 2, 9, 98, 4, 98, 1, 3, 4, 9, 1, 98, 98, 4, 2, 3, 98, 98, 1, 99, 9, 98, 98, 3, 98, 98, 4, 98, 2, 98, 4, 2, 1, 1, 9, 2, 4}))
	// fmt.Println(pickingNumbers2([]int32{9, 6, 13, 16, 5, 18, 4, 10, 3, 19, 4, 5, 8, 1, 13, 10, 20, 17, 15, 10, 6, 10, 13, 20, 18, 17, 7, 10, 6, 5, 16, 18, 13, 20, 19, 7, 16, 13, 20, 17, 4, 17, 8, 19, 12, 7, 17, 1, 18, 3, 16, 4, 5, 3, 15, 17, 6, 17, 14, 11, 11, 7, 11, 6, 15, 15, 12, 6, 17, 19, 8, 6, 13, 9, 10, 19, 14, 18, 7, 9, 11, 16, 11, 20, 4, 20, 10, 7, 8, 4, 2, 12, 11, 8, 12, 13, 19, 8, 8, 5}))
	// fmt.Println(pickingNumbers2([]int32{7, 12, 13, 19, 17, 7, 3, 18, 9, 18, 13, 12, 3, 13, 7, 9, 18, 9, 18, 9, 13, 18, 13, 13, 18, 18, 17, 17, 13, 3, 12, 13, 19, 17, 19, 12, 18, 13, 7, 3, 3, 12, 7, 13, 7, 3, 17, 9, 13, 13, 13, 12, 18, 18, 9, 7, 19, 17, 13, 18, 19, 9, 18, 18, 18, 19, 17, 7, 12, 3, 13, 19, 12, 3, 9, 13, 19, 12, 18, 13, 18, 18, 18, 17, 13, 3, 18, 19, 7, 12, 9, 18, 3, 13, 13, 9, 7}))

}
