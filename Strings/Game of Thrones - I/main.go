package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s := readline(r)
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	odd := 0
	for _, v := range m {
		if v%2 == 1 {
			odd++
		}
		if odd > 1 {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}

func readline(r *bufio.Reader) (s string) {
	for {
		line, prefix, _ := r.ReadLine()
		s = s + string(line)
		if !prefix {
			break
		}
	}
	return
}
