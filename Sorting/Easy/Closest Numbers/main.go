package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	_ = n
	var inputs []int
	for _, v := range strings.Fields(readline(r)) {
		num, _ := strconv.Atoi(v)
		inputs = append(inputs, num)
	}
	print(closestNumbers(inputs))
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

func closestNumbers(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	out := make([]int, 2)
	out[0], out[1] = arr[0], arr[1]
	diff := arr[1] - arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] > diff {
			continue
		}
		if arr[i+1]-arr[i] < diff {
			out = out[:0]
			diff = arr[i+1] - arr[i]
		}
		out = append(out, arr[i], arr[i+1])
	}
	return out
}

func print(a []int) {
	for _, v := range a {
		fmt.Printf("%d ", v)
	}
	fmt.Println()
}
