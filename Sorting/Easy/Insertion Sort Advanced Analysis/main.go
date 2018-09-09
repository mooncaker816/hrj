package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var m, num int32
		fmt.Scanf("%d\n", &m)
		inputs := make([]int32, m)
		for j := 0; j < int(m); j++ {
			fmt.Scanf("%d", &num)
			inputs[j] = num
		}
		fmt.Scanln()
		fmt.Println(insertionSort(inputs))
	}
}

func insertionSort(arr []int32) int64 {
	_, shifts := mergeSort(arr)
	return shifts
}

func mergeSort(a []int32) ([]int32, int64) {
	if len(a) <= 1 {
		return a, 0
	}
	mid := len(a) / 2
	left, cntLeft := mergeSort(a[:mid])
	right, cntRight := mergeSort(a[mid:])
	out, cnt := merge(left, right)
	return out, cntLeft + cntRight + cnt
}

func merge(a, b []int32) ([]int32, int64) {
	c := make([]int32, 0, len(a)+len(b))
	cnt := int64(0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			c = append(c, a[i])
			i++
		} else {
			c = append(c, b[j])
			cnt += int64(len(a) - i)
			j++
		}
	}
	c = append(c, a[i:]...)
	c = append(c, b[j:]...)
	return c, cnt
}
