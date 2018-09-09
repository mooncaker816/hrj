package main

import "fmt"

func main() {
	var n, v int32
	fmt.Scanf("%d\n", &v)
	fmt.Scanf("%d\n", &n)
	inputs := make([]int32, n)
	for i := 0; i < int(n); i++ {
		var num int32
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}
	fmt.Println(introTutorial(v, inputs))
}

func introTutorial(v int32, a []int32) int32 {
	for i := range a {
		if v <= a[i] {
			return int32(i)
		}
	}
	return int32(len(a))
}
