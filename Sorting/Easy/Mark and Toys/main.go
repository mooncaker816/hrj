package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)
	prices := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		prices[i] = num
	}
	fmt.Println(maximumToys(prices, k))
}

func maximumToys(prices []int, k int) int {
	sort.Ints(prices)
	i, sum := 0, 0
	for {
		sum += prices[i]
		if sum > k {
			break
		}
		i++
	}
	return i
}
