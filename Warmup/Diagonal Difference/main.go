package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type matrix [][]string

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	in1 := r.Text()
	size, err := strconv.Atoi(in1)
	if err != nil {
		fmt.Println(err)
	}
	var m matrix
	for r.Scan() {
		line := r.Text()
		if line != "" {
			s := strings.Fields(string(line))
			m = append(m, s)
		}
	}
	sum1, sum2 := 0, 0
	for i, v1 := range m {
		for j, v2 := range v1 {
			if i == j {
				val1, _ := strconv.Atoi(v2)
				sum1 += val1
			}
			if i+j == size-1 {
				val2, _ := strconv.Atoi(v2)
				sum2 += val2
			}
		}
	}
	ret := sum1 - sum2
	if ret < 0 {
		ret = -ret
	}
	fmt.Println(ret)
}
