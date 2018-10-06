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
	strs := strings.Fields(readline(r))
	inputs := make([]int, n)
	for i := 0; i < n; i++ {
		inputs[i], _ = strconv.Atoi(strs[i])
	}
	fmt.Printf("Array is sorted in %d swaps.\n", countSwaps(inputs))
	fmt.Printf("First Element: %d\n", inputs[0])
	fmt.Printf("Last Element: %d\n", inputs[n-1])
}

func readline(r *bufio.Reader) (s string) {
	for {
		line, prefix, _ := r.ReadLine()
		s = s + string(line)
		if !prefix {
			break
		}
	}
	return strings.TrimSpace(s)
}

func countSwaps(a []int) int64 {
	cnt := int64(0)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-i-1; j++ {
			// Swap adjacent elements if they are in decreasing order
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
				cnt++
			}
		}
	}
	return cnt
}
