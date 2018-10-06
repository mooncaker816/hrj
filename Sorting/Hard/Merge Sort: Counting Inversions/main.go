package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/ctci-merge-sort/problem?h_l=interview&playlist_slugs%5B%5D=interview-preparation-kit&playlist_slugs%5B%5D=sorting

func main() {
	r := bufio.NewReader(os.Stdin)
	d, _ := strconv.Atoi(readline(r))
	for i := 0; i < d; i++ {
		n, _ := strconv.Atoi(readline(r))
		strs := strings.Fields(readline(r))
		inputs := make([]int, n)
		for j := 0; j < n; j++ {
			inputs[j], _ = strconv.Atoi(strs[j])
		}
		fmt.Println(countInversions(inputs))
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
	return strings.TrimSpace(s)
}

func countInversions(a []int) int64 {
	_, cnt := mergeSort(a)
	return cnt
}

func mergeSort(a []int) ([]int, int64) {
	mid := len(a) / 2
	if mid == 0 {
		return a, 0
	}
	left, cnt0 := mergeSort(a[:mid])
	right, cnt1 := mergeSort(a[mid:])
	b, cnt := merge(left, right)
	return b, cnt0 + cnt1 + cnt
}

func merge(a, b []int) ([]int, int64) {
	out := make([]int, 0, len(a)+len(b))
	cnt := int64(0)
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
			cnt += int64(len(a) - i)
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out, cnt
}
