package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	weights := make([]int, 0, 100000000)
	for i := 0; i < len(s); i++ {
		w := int(rune(s[i]) - 'a' + 1)
		weights = append(weights, w)
		for j := i + 1; j < len(s); j++ {
			if s[i] != s[j] {
				i = j - 1
				break
			}
			weights = append(weights, w*(j-i+1))
		}
	}
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scanf("%d\n", &num)
		result := "No"
		for _, v := range weights {
			if num == v {
				result = "Yes"
				break
			}
		}
		fmt.Println(result)
	}
}
