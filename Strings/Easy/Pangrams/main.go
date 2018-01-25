package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _, _ := r.ReadLine()
	s := string(line)
	s0 := "abcdefghijklmnopqrstuvwxyz"
	m := make(map[rune]int)
	for _, r := range s0 {
		m[r] = 0
	}
	s = strings.ToLower(s)
	for _, r := range s {
		if _, ok := m[r]; ok {
			m[r]++
		}
	}
	result := "pangram"
	for _, v := range m {
		if v == 0 {
			result = "not pangram"
			break
		}
	}
	fmt.Println(result)
}
