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
		seen := make(map[rune]bool)
		cost := 0
		for _, r := range s {
			if seen[r] {
				continue
			}
			cost++
			seen[r] = true
		}
		fmt.Println(cost)
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
