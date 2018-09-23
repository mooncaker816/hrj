package main

import "fmt"

func main() {
	var s, t, a, b, m, n, cntA, cntB int
	fmt.Scanf("%d %d\n", &s, &t)
	fmt.Scanf("%d %d\n", &a, &b)
	fmt.Scanf("%d %d\n", &m, &n)
	for i := 0; i < m; i++ {
		var x int
		fmt.Scanf("%d", &x)
		if a+x >= s && a+x <= t {
			cntA++
		}
	}
	for i := 0; i < n; i++ {
		var y int
		fmt.Scanf("%d", &y)
		if b+y >= s && b+y <= t {
			cntB++
		}
	}
	fmt.Println(cntA)
	fmt.Println(cntB)
}
