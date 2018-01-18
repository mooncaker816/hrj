package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	size, _ := strconv.Atoi(r.Text())
	r.Scan()
	s := strings.Fields(r.Text())
	p, n := 0, 0
	for _, v := range s {
		val, _ := strconv.Atoi(v)
		switch {
		case val > 0:
			p++
		case val < 0:
			n++
		}
	}
	ret1 := float64(p) / float64(size)
	ret2 := float64(n) / float64(size)
	ret3 := float64(size-p-n) / float64(size)
	fmt.Printf("%.6f\n", ret1)
	fmt.Printf("%.6f\n", ret2)
	fmt.Printf("%.6f\n", ret3)

}
