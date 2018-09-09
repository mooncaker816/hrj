package main

import "fmt"

func main() {
	var n int32
	fmt.Scanf("%d\n", &n)
	inputs := make([]int32, n)
	for i := 0; i < int(n); i++ {
		var num int32
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}
	fmt.Println(runningTime(inputs))
}

func runningTime(a []int32) int32 {
	var cnt int32
	for i := 1; i < len(a); i++ {
		for j := i; j >= 1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
				cnt++
			} else {
				break
			}
		}
	}
	return cnt
}
