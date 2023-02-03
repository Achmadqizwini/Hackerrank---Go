package main

import "fmt"

func viralAdvertising(n int32) int32 {
	var share int32 = 5
	var like, cumulative int32
	for i := int32(1); i < n; i++ {
		like = int32(share) / 2
		cumulative += like
		share = like * 3
	}
	return cumulative
}

func main() {
	fmt.Println(viralAdvertising(3))
	fmt.Println(viralAdvertising(4))
	fmt.Println(viralAdvertising(5))
}
