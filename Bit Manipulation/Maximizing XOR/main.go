package main

import (
	"fmt"
	"math/bits"
)

const wordlen = 32 << (^uint(0) >> 63)

func main() {
	var l, r, result uint
	fmt.Scanf("%d\n", &l)
	fmt.Scanf("%d\n", &r)
	// 朴素版本
	// for i := l; i <= r; i++ {
	// 	for j := i; j <= r; j++ {
	// 		temp := i ^ j
	// 		if temp > result {
	// 			result = temp
	// 		}
	// 	}
	// }

	// rshaghoulian
	// To maximize A xor B, we want A and B to differ as much as possible at every bit index.

	// We first find the most significant bit that we can force to differ by looking at L and R.

	// For all of the lesser significant bits in A and B, we can always ensure that they differ and still have L <= A <= B <= R.

	// Our final answer will be the number represented by all 1s starting from the most significant bit that differs between A and B

	// L = 10111
	// R = 11100
	// 	_X___  <-- that's most significant differing bit
	// 	01111  <-- here's our final answer
	xord := l ^ r
	significantBit := bits.UintSize - bits.LeadingZeros(xord) - 1
	//significantBit := wordlen - bits.LeadingZeros(xord) - 1
	result = (1 << (uint(significantBit) + 1)) - 1
	fmt.Println(result)
}
