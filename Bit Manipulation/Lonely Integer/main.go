package main

import "fmt"

func main() {
	var n, result int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		result ^= num
	}
	fmt.Println(result)
}
