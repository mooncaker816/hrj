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
		num := 0
		for j, k := 0, len(s)-1; j <= k; j, k = j+1, k-1 {
			num += changeToSame(s[j], s[k])
		}
		fmt.Println(num)
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

func changeToSame(a, b byte) int {
	if a == b {
		return 0
	}
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}
