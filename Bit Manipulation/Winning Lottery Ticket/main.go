package main

import (
	"fmt"
	"strconv"
)

func main() {
	var q int
	fmt.Scanf("%d\n", &q)
	tickets := make(map[uint]int)
	for i := 0; i < q; i++ {
		var n string
		fmt.Scanf("%s\n", &n)
		tickets[bitmap(n)]++
	}
	var total, left int
	left = q
	seen := make(map[uint]bool)
	for k0, v0 := range tickets {
		seen[k0] = true
		left -= v0
		if k0 == 1<<10-1 {
			total += v0*left + v0*(v0-1)/2
			continue
		}
		for k1, v1 := range tickets {
			_, ok := seen[k1]
			if !ok && (k0|k1 == 1<<10-1) {
				total += v0 * v1
			}
		}
	}
	fmt.Println(total)
}

func bitmap(n string) uint {
	var res uint
	for _, r := range n {
		digit, _ := strconv.Atoi(string(r))
		res |= 1 << uint(digit)
		if res == 1<<10-1 {
			break
		}
	}
	return res
}
