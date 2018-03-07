package main

import (
	"fmt"
	"math"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n uint32
		fmt.Scanf("%d\n", &n)
		fmt.Println(math.MaxUint32 - n)
	}
}
