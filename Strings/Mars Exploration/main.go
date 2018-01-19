package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	num := 0
	for i := 0; i <= len(s)-3; i = i + 3 {
		if s[i] != 'S' {
			num++
		}
		if s[i+1] != 'O' {
			num++
		}
		if s[i+2] != 'S' {
			num++
		}
	}
	fmt.Println(num)
}
