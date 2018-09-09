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
	inputs := make([][]string, n)
	for i := 0; i < n; i++ {
		inputs[i] = strings.Fields(readline(r))
	}
	countSort(inputs)
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

func countSort(arr [][]string) {
	k := 100
	counts := make([]int, k)
	keys := make([]int, len(arr))
	for i, v := range arr {
		if i < len(arr)/2 {
			arr[i][1] = "-"
		}
		num, _ := strconv.Atoi(v[0])
		keys[i] = num
		counts[num]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	b := make([]string, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		b[counts[keys[i]]-1] = arr[i][1]
		counts[keys[i]]--
	}
	for _, v := range b {
		fmt.Printf("%s ", v)
	}
	fmt.Println()
}
