package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int32
	fmt.Scanf("%d %d\n", &n, &k)
	inputs := make([]int32, n)
	for i := 0; i < int(n); i++ {
		var num int32
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}

	fmt.Println(hackerlandRadioTransmitters(inputs, k))
}

func hackerlandRadioTransmitters(a []int32, k int32) int32 {
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	var loc, cnt int32
	i := 0
	for i < len(a) {
		cnt++
		loc = a[i] + k
		for ; i < len(a); i++ {
			if a[i] > loc {
				break
			}
		}

		i--
		loc = a[i] + k

		for ; i < len(a); i++ {
			if a[i] > loc {
				break
			}
		}
	}
	return cnt
}
