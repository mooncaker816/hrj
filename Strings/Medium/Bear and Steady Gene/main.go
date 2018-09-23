package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	var s string
	fmt.Scanf("%s\n", &s)
	fmt.Println(steadyGene(s))
}

func steadyGene(gene string) int {
	m := make(map[byte]int, 4)
	total := len(gene)
	for i := 0; i < len(gene); i++ {
		m[gene[i]]++
	}
	i, j := 0, 0
	minV := total
	for ; j < len(gene) && i <= j; j++ {
		m[gene[j]]--
		for chkMap(m, total) {
			minV = min(minV, j-i+1)
			m[gene[i]]++
			i++
		}
	}
	return minV
}

func chkMap(m map[byte]int, size int) bool {
	for _, v := range m {
		if v > size/4 {
			return false
		}
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
