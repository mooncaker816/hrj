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
	r.ReadString('\n')
	in2, _ := r.ReadString('\n')
	s := strings.Fields(string(in2))
	var sum int64
	exp := ""
	for i, v := range s {
		vi, _ := strconv.ParseInt(v, 10, 64)
		sum += vi
		if i > 0 {
			exp += "+"
		}
		exp += v
	}
	//fmt.Printf("%s = %d\n", exp, sum)
	fmt.Println(sum)
}
