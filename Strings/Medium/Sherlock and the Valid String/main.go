package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s := readline(r)
	m := make(map[rune]int)
	for _, r := range s {
		m[r]++
	}
	occurs := make([]int, 0)
	for _, v := range m {
		occurs = append(occurs, v)
	}
	sort.Ints(occurs)
	result := "YES"
	// if len(occurs) == 2 {
	// 	if occurs[0] != occurs[1] && occurs[0]+1 != occurs[1] {
	// 		if occurs[0] > 1 {
	// 			result = "NO"
	// 		}
	// 	}
	// }
	// if len(occurs) > 2 {
	if occurs[0] == 1 {
		if len(occurs) > 1 {
			if occurs[1] != occurs[len(occurs)-1] {
				result = "NO"
			}
		}
		// for i := 1; i < len(occurs)-1; i++ {
		// 	if occurs[i] != occurs[i+1] {
		// 		result = "NO"
		// 		break
		// 	}
		// }
	}
	if occurs[0] > 1 {
		if len(occurs) > 2 {
			if occurs[0] != occurs[len(occurs)-2] {
				result = "NO"
			} else {
				if occurs[0] != occurs[len(occurs)-1] && occurs[0]+1 != occurs[len(occurs)-1] {
					result = "NO"
				}
			}
		}
		if len(occurs) == 2 {
			if occurs[0] != occurs[1] && occurs[0]+1 != occurs[1] {
				result = "NO"
			}
		}
		// for i := 0; i < len(occurs)-1; i++ {
		// 	if occurs[i] != occurs[i+1] {
		// 		if i < len(occurs)-2 {
		// 			result = "NO"
		// 			break
		// 		}
		// 		if occurs[i]+1 != occurs[i+1] {
		// 			result = "NO"
		// 			break
		// 		}
		// 	}
		// }
	}
	// }
	fmt.Println(result)
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
