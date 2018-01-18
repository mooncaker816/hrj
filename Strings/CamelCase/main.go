package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanf("%s\n", &s)
	num := 1
	for _, v := range s {
		if strings.ToUpper(string(v)) == string(v) {
			num++
		}
	}
	fmt.Println(num)
}
