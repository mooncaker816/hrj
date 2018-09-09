package main

import (
	"bufio"
	"fmt"
)

func main() {
	var n int32
	fmt.Scanf("%d\n", &n)
	inputs := make([]int32, n)
	for i := 0; i < int(n); i++ {
		var num int32
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}
	insertionSort1(n, inputs)
}

func readline(r *bufio.Reader) (s string) {
	for {
		line, prefix, _ := r.ReadLine()
		s = s + string(line)
		if !prefix {
			break
		}
	}
	return
}

func insertionSort1(n int32, a []int32) {
	v := a[n-1]
	// inserted := false
	j := n - 2
	for ; j >= 0; j-- {
		if v < a[j] {
			a[j+1] = a[j]
			print(a)
			continue
		}
		// a[j+1] = v
		// inserted = true
		// print(a)
		break
	}
	a[j+1] = v
	print(a)

	// if !inserted {
	// 	a[0] = v
	// 	print(a)
	// }
}

func print(a []int32) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
