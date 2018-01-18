package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	r := bufio.NewScanner(os.Stdin)
	r.Scan()
	size, _ := strconv.Atoi(r.Text())
	for i := 0; i < size; i++ {
		for snum := size - i - 1; snum > 0; snum-- {
			fmt.Print(" ")
		}
		for num := i + 1; num > 0; num-- {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
