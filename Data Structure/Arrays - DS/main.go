package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	s := strings.Fields(readline(r))
	for i := n - 1; i >= 0; i-- {
		fmt.Printf("%s ", s[i])
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
