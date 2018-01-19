package main

import (
	"fmt"
)

type chars map[rune][]int

func main() {
	var n int
	var s string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &s)
	m := make(map[rune][]int, 0)
	for i, v := range s {
		if _, ok := m[v]; ok {
			m[v] = append(m[v], i)
		} else {
			m[v] = []int{i}
		}
	}
	num := 0
	for _, r := range s {
		for k, v := range m {
			if k == r {
				continue
			}
			if compare(m[r], v) {
				//fmt.Printf("%s\n%s\n%v\n%v\n", string(k), string(r), m[r], v)
				if len(m[r])+len(v) > num {
					num = len(m[r]) + len(v)
				}
			}
		}
	}
	fmt.Printf("%d\n", num)
}

func compare(s1, s2 []int) bool {
	if len(s1)-len(s2) > 1 || len(s1)-len(s2) < -1 {
		return false
	}
	if len(s1) == len(s2) {
		if s1[0] < s2[0] {
			for i := 0; i < len(s1)-1; i++ {
				if s1[i] >= s2[i] || s1[i+1] <= s2[i] {
					return false
				}
			}
			if s1[len(s1)-1] >= s2[len(s1)-1] {
				return false
			}
		}
		if s1[0] > s2[0] {
			for i := 0; i < len(s1)-1; i++ {
				if s2[i] >= s1[i] || s2[i+1] <= s1[i] {
					return false
				}
			}
			if s2[len(s1)-1] >= s1[len(s1)-1] {
				return false
			}
		}
	}
	if len(s1) == len(s2)+1 {
		if s1[0] >= s2[0] || s2[0] >= s1[1] {
			return false
		}
		for i := 0; i < len(s2)-1; i++ {
			if s2[i] >= s1[i+1] || s2[i+1] <= s1[i+1] {
				return false
			}
		}
		if s2[len(s2)-1] >= s1[len(s2)] {
			return false
		}
	}
	if len(s2) == len(s1)+1 {
		if s2[0] >= s1[0] || s1[0] >= s2[1] {
			return false
		}
		for i := 0; i < len(s1)-1; i++ {
			if s1[i] >= s2[i+1] || s1[i+1] <= s2[i+1] {
				return false
			}
		}
		if s1[len(s1)-1] >= s2[len(s1)] {
			return false
		}
	}
	return true
}
