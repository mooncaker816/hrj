package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	max, min := 0, 0
	maxCnt, minCnt := 0, 0
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		if i == 0 {
			max, min = num, num
		} else {
			if num > max {
				max = num
				maxCnt++
			} else if num < min {
				min = num
				minCnt++
			}
		}
	}
	fmt.Println(maxCnt, minCnt)
}
