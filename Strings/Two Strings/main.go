package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	for i := 0; i < n; i++ {
		s1 := readline(r)
		s2 := readline(r)
		result := "NO"
		for i := 0; i < len(s1); i++ {
			if strings.Contains(s2, s1[i:i+1]) {
				result = "YES"
				break
			}
		}
		fmt.Println(result)
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
