package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	for i := 0; i < n; i++ {
		s := readline(r)
		if len(s)%2 == 1 {
			fmt.Println("-1")
			continue
		}
		s1 := s[:len(s)/2]
		s2 := s[len(s)/2:]
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
		}
		fmt.Println(delta)
	}
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
