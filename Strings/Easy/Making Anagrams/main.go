package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s1 := readline(r)
	s2 := readline(r)
	m1 := make(map[rune]int, 26)
	m2 := make(map[rune]int, 26)
	for i := 0; i < 26; i++ {
		c := 'a' + i
		m1[rune(c)] = 0
		m2[rune(c)] = 0
	}
	for _, r := range s1 {
		m1[r]++
	}
	for _, r := range s2 {
		m2[r]++
	}
	delta := 0
	for r1, v1 := range m1 {
		if (v1 - m2[r1]) > 0 {
			delta += v1 - m2[r1]
		}
		if (v1 - m2[r1]) < 0 {
			delta += m2[r1] - v1
		}
	}
	fmt.Println(delta)
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
