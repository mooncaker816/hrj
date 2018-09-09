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
	print(countingSort(inputs))
}

func countingSort(arr []int32) []int32 {
	out := make([]int32, 100)
	for i := 0; i < len(arr); i++ {
		out[arr[i]]++
	}
	return out
}

func print(a []int32) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
