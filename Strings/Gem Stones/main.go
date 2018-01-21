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
	prev := ^uint32(0)
	for i := 0; i < n; i++ {
		s := readline(r)
		var this uint32
		for _, r := range s {
			this = this | 1<<uint32(r-'a')
		}
		prev = prev & this
	}
	count := 0
	for prev != 0 {
		count++
		prev = prev & (prev - 1)
	}
	fmt.Println(count)
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
