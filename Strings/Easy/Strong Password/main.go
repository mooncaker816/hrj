package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	var password string
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%s\n", &password)
	numbers := "0123456789"
	lower_case := "abcdefghijklmnopqrstuvwxyz"
	upper_case := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	special_characters := "!@#$%^&*()-+"
	must := 4
	numfound, lowfound, upfound, specialfound := false, false, false, false
	for _, r := range password {
		if strings.IndexRune(numbers, r) != -1 && !numfound {
			must--
			numfound = true
			continue
		}
		if strings.IndexRune(lower_case, r) != -1 && !lowfound {
			must--
			lowfound = true
			continue
		}
		if strings.IndexRune(upper_case, r) != -1 && !upfound {
			must--
			upfound = true
			continue
		}
		if strings.IndexRune(special_characters, r) != -1 && !specialfound {
			must--
			specialfound = true
			continue
		}
		if must == 0 {
			break
		}
	}
	if must < 6-n {
		if 6-n < 0 {
			fmt.Println(0)
		} else {
			fmt.Println(6 - n)
		}
	} else {
		fmt.Println(must)
	}
}
