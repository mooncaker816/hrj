package main

import (
	"fmt"
)

func main() {
	var n, k int
	var s string
	fmt.Scanf("%d %d\n", &n, &k)
	fmt.Scanf("%s\n", &s)
	fmt.Println(highestValuePalindrome(s, n, k))
}

func highestValuePalindrome(s string, n int, k int) string {
	return string(genHighestP(preLoad(s), n))
}

func preLoad(s string) []int {
	stat := make([]int, 10)
	for i := 0; i < len(s); i++ {
		stat[s[i]-'0']++
	}
	return stat
}

func genHighestP(stat []int, size int) []byte {
	p := make([]byte, size)
	oddCnt := 0
	if size&1 == 1 {
		oddCnt--
	}
	var odd byte
	pos := 0
	for i := len(stat) - 1; i >= 0; i-- {
		c := byte(i&0xf | 0x30) // decimal value to digit char
		if stat[i]&1 == 1 {
			oddCnt++
			if oddCnt > 0 {
				return nil
			}
			odd = c
		}
		for j := 0; j < stat[i]/2; j++ {
			p[pos] = c
			p[size-pos-1] = c
			pos++
		}
	}
	if odd != 0 {
		p[size/2] = odd
	}
	return p
}
