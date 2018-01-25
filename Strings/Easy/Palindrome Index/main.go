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
		idx := -1
		for j, k := 0, len(s)-1; j < len(s); j, k = j+1, k-1 {
			if s[j] == s[k] {
				continue
			}
			if s[j+1] == s[k] {
				ss := s[j+2 : k]
				if isPalindrome(ss) {
					idx = j
					break
				}
			}
			if s[j] == s[k-1] {
				ss := s[j+1 : k-1]
				if isPalindrome(ss) {
					idx = k
					break
				}
			}
			break
		}
		fmt.Println(idx)
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

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
