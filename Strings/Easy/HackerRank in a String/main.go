package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	p := "hackerrank"
	for i := 0; i < n; i++ {
		var q string
		fmt.Scanf("%s\n", &q)
		m := build(q)
		prev := -1
		result := "YES"
		for _, r := range p {
			tmp := checkpath(r, prev, m)
			if tmp < 0 {
				result = "NO"
				break
			}
			prev = tmp
		}
		fmt.Println(result)
	}
}

type status struct {
	seq []int
	pos int
}

func build(s string) map[rune]status {
	m := make(map[rune]status, 0)
	for i, r := range s {
		if v, ok := m[r]; ok {
			v.seq = append(v.seq, i)
			m[r] = v
		} else {
			m[r] = status{[]int{i}, 0}
		}
	}
	return m
}

func checkpath(r rune, prev int, m map[rune]status) int {
	if st, ok := m[r]; ok {
		for st.pos < len(st.seq) {
			if st.seq[st.pos] > prev {
				m[r] = st
				return st.seq[st.pos]
			}
			st.pos++
		}

	}
	return -1
}
