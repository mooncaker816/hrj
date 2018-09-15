package main

import (
	"fmt"
)

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Scanf("%s\n", &s)
	p := highestValuePalindrome(s, n, k)
	if p == "" {
		fmt.Println(-1)
	} else {
		fmt.Println(p)
	}
}

func highestValuePalindrome(s string, n int, k int) string {
	b := []byte(s)
	cnt, path := leastStepsToP(b)
	switch {
	case cnt > k:
		return ""
	case cnt < k:
		additionUpdate(b, path, k-cnt)
	}
	return string(b)
}

func leastStepsToP(b []byte) (cnt int, path map[int]struct{}) {
	path = make(map[int]struct{})
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		switch {
		case b[i] > b[j]:
			b[j] = b[i]
			cnt++
			path[i] = struct{}{}
		case b[i] < b[j]:
			b[i] = b[j]
			cnt++
			path[i] = struct{}{}
		}
	}
	return cnt, path
}

func additionUpdate(b []byte, path map[int]struct{}, cnt int) {
	for i, j, k := 0, len(b)-1, 0; i <= j && cnt > 0; i, j = i+1, j-1 {
		if b[i] < '9' {
			if _, ok := path[i]; ok || i == j {
				b[i] = '9'
				b[j] = '9'
				cnt--
				k++
			} else {
				if cnt > 1 {
					b[i] = '9'
					b[j] = '9'
					cnt -= 2
				}
			}
		}
	}
}
