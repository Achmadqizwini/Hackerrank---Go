package main

import "fmt"

func howManyGames(p int32, d int32, m int32, s int32) int32 {
	var res int32 = 0
	if s < p {
		return 0
	}
	for p > m {
		s -= p
		res++
		p -= d
		if s < p {
			break
		}
		if p < m {
			p = m
		}
	}
	if s > p {
		a := s / p
		res += a
	}
	return res
}

func main() {
	fmt.Println(howManyGames(20, 3, 6, 80))
	fmt.Println(howManyGames(20, 3, 6, 85))
	fmt.Println(howManyGames(100, 19, 1, 180))
}
