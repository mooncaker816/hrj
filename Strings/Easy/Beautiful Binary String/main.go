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
	strconv.Atoi(readline(r))
	s := readline(r)
	rep := strings.NewReplacer("010", "")
	fmt.Println((len(s) - len(rep.Replace(s))) / 3)
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
