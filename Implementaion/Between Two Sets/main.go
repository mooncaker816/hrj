package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d\n", &n, &m)
	aMax := 0
	a := make([]int, n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d", &num)
		a[i] = num
		if num > aMax {
			aMax = num
		}
	}
	bMin := 100
	b := make([]int, m)
	for i := 0; i < m; i++ {
		var num int
		fmt.Scanf("%d", &num)
		b[i] = num
		if num < bMin {
			bMin = num
		}
	}
	cnt := 0
	for i := aMax; i <= bMin; i++ {
		ok := true
		for j := 0; j < n; j++ {
			if i%a[j] != 0 {
				ok = false
				break
			}
		}
		if ok {
			for j := 0; j < m; j++ {
				if b[j]%i != 0 {
					ok = false
					break
				}
			}
		}
		if ok {
			cnt++
		}
	}
	fmt.Println(cnt)
}
