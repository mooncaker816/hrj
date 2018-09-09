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
	print(quickSort(inputs))
}

func quickSort(arr []int32) []int32 {
	left := make([]int32, 0)
	right := make([]int32, 0)
	equal := make([]int32, 1)
	equal[0] = arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] == arr[0] {
			equal = append(equal, arr[i])
		} else if arr[i] > arr[0] {
			right = append(right, arr[i])
		} else {
			left = append(left, arr[i])
		}
	}
	left = append(left, equal...)
	left = append(left, right...)
	return left
}

func print(a []int32) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
