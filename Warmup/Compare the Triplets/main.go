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
	in1, _ := r.ReadString('\n')
	in2, _ := r.ReadString('\n')
	s1 := strings.Fields(string(in1))
	s2 := strings.Fields(string(in2))
	p1, p2 := 0, 0
	for i := 0; i < len(s1); i++ {
		num1, _ := strconv.Atoi(s1[i])
		num2, _ := strconv.Atoi(s2[i])
		if num1 > num2 {
			p1++
		}
		if num1 < num2 {
			p2++
		}
	}
	fmt.Println(p1, p2)
}
