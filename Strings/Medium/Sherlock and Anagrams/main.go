package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Scanf("%s", &s)
		fmt.Println(sherlockAndAnagrams(s))
	}
}

func sherlockAndAnagrams(s string) int {
	cnt := 0
	stat := make(map[string]int)

	// 遍历每一个子串
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subStr := sortString(s[i:j]) // 如果两个子串有同样的字符，那么排序后肯定相同
			stat[subStr]++               // 每一组拥有相同字符的子串的个数
		}
	}
	for _, n := range stat {
		cnt += n * (n - 1) / 2 // 从n个相同字符的子串中取出两个的取法 = C(n,2)
	}
	return cnt
}

func sortString(s string) string {
	a := []byte(s)
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return string(a)
}
