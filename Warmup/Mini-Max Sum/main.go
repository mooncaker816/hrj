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
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	var s sort.StringSlice
	s = strings.Fields(r.Text())
	sort.Sort(s)
	var min, max int64
	for i, v := range s {
		val, _ := strconv.ParseInt(v, 10, 64)
		if i < len(s)-1 {
			min += val
		}
		if i > 0 {
			max += val
		}
	}
	fmt.Println(min, max)
}
