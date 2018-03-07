package main

import (
	"fmt"
	"math/bits"
)

func main() {
	var q int
	fmt.Scanf("%d\n", &q)
	for i := 0; i < q; i++ {
		var n uint
		fmt.Scanf("%d\n", &n)
		fmt.Println(1<<uint(bits.Len(n)) - 1 - n)
	}
}
