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
	insertionSort2(n, inputs)
}

func insertionSort2(n int32, a []int32) {
	for i := 1; i < int(n); i++ {
		for j := i; j >= 1; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
		print(a)
	}
}

func print(a []int32) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
