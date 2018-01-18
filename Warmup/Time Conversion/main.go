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
	timeStr := r.Text()
	hh := timeStr[:2]
	mm := timeStr[3:5]
	ss := timeStr[6:8]
	ampm := timeStr[8:]
	if ampm == "PM" && hh != "12" {
		h, _ := strconv.Atoi(hh)
		hh = strconv.Itoa(h + 12)
	}
	if ampm == "AM" && hh == "12" {
		hh = "00"
	}
	out := hh + ":" + mm + ":" + ss

	fmt.Printf("%s\n", out)

}
