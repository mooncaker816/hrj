package main

import (
	"fmt"
)

func main() {
	var n, key int
	var s string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &s)
	fmt.Scanf("%d\n", &key)
	var nr []rune
	for _, r := range s {
		if r <= 'z' && r >= 'a' {
			nr = append(nr, rune((int(r-'a')+key)%26+'a'))
			continue
		}
		if r <= 'Z' && r >= 'A' {
			nr = append(nr, rune((int(r-'A')+key)%26+'A'))
			continue
		}
		nr = append(nr, r)
	}
	fmt.Println(string(nr))
}
