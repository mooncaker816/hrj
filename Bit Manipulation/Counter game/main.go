package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n uint64
		fmt.Scanf("%d\n", &n)
		//counter := bits.OnesCount64(n) + bits.TrailingZeros64(n) - 1
		counter := bits.OnesCount64(n - 1)
		if counter&1 == 1 {
			fmt.Println("Louise")
		} else {
			fmt.Println("Richard")
		}
	}
}
