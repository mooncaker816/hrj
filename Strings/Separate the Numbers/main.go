package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	for i := 0; i < n; i++ {
		var s string
		var ret int
		fmt.Scanf("%s\n", &s)
		if len(s) == 1 || s[0] == '0' {
			fmt.Println("NO")
			continue
		}
		maxstep := len(s) / 2
		for step := 1; step <= maxstep; step++ {
			ret = check(step, s)
			if ret > 0 {
				break
			}
		}
		if ret > 0 {
			fmt.Printf("YES %d\n", ret)
		} else {
			fmt.Println("NO")
		}
	}
}

func check(step int, s string) int {
	//startval := strconv.Atoi(s[0 : step])
	val, _ := strconv.Atoi(s[0:step])
	ret := val
	left := s[step:]
	for {
		val++
		next := strconv.Itoa(val)
		step = len(next)
		if strings.HasPrefix(left, next) {
			left = left[step:]
			if len(left) == 0 {
				break
			}
		} else {
			return -1
		}
	}
	return ret
}
