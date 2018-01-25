package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	line, _, _ := r.ReadLine()
	n, _ := strconv.Atoi(string(line))
	for i := 0; i < n; i++ {
		var s string
		prex := true
		for prex {
			line, prex, _ = r.ReadLine()
			s += string(line)
		}
		ret := "Funny"
		for j, k := 0, len(s)-1; j <= len(s)/2; j, k = j+1, k-1 {
			if s[j]-s[j+1] != s[k]-s[k-1] && s[j]-s[j+1] != s[k-1]-s[k] {
				ret = "Not Funny"
				break
			}
		}
		fmt.Println(ret)
	}
}
