package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	inputs := make([]string, n)
	for i := 0; i < n; i++ {
		inputs[i] = readline(r)
	}
	bigSorting(inputs)
	for _, v := range inputs {
		fmt.Println(v)
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

func bigSorting(unsorted []string) []string {
	sort.Slice(unsorted, func(i, j int) bool {
		if len(unsorted[i]) != len(unsorted[j]) {
			return len(unsorted[i]) < len(unsorted[j])
		}
		return unsorted[i] < unsorted[j]
	})
	return unsorted
}
