package main

import (
	"fmt"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	for i := 0; i < t; i++ {
		var n int
		var res uint
		fmt.Scanf("%d\n", &n)
		for j := 0; j < n; j++ {
			var m uint
			fmt.Scanf("%d", &m)
			if (n-j)*(j+1)&1 == 1 {
				res ^= m
			}
		}
		fmt.Println(res)
	}
}
