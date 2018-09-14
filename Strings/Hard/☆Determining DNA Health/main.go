package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// case0 : 0 19
// case1 : 15806635 20688978289
// case2 : 0 3120724375
// case3 : 3218660 11137051
// case4 : 207177686 3128254990

type standard struct {
	start, end int
	value      string
}

func main() {
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()
	startTime := time.Now()
	_ = startTime
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	_ = n
	genes := strings.Fields(readline(r))
	healthStr := strings.Fields(readline(r))
	// health := make([]int, len(healthStr))
	// for i, hs := range healthStr {
	// 	h, _ := strconv.Atoi(hs)
	// 	health[i] = h
	// }

	healthMap := make(map[string][][2]int, n)
	for i, hs := range healthStr {
		h, _ := strconv.Atoi(hs)
		// if healthMap[genes[i]] == nil {
		// 	healthMap[genes[i]] = make([]int, n)
		// }
		// healthMap[genes[i]][i] = h
		if healthMap[genes[i]] == nil {
			healthMap[genes[i]] = make([][2]int, 0)
		}
		healthMap[genes[i]] = append(healthMap[genes[i]], [2]int{i, h})
	}
	// fmt.Println(len(healthMap))
	// var wg sync.WaitGroup
	// for _, counts := range healthMap {
	// 	wg.Add(1)
	// 	go func(counts []int) {
	// 		for i := 1; i < len(counts); i++ {
	// 			counts[i] += counts[i-1]
	// 		}
	// 		wg.Done()
	// 	}(counts)
	// }
	// wg.Wait()
	m, _ := strconv.Atoi(readline(r))
	// standards := make([]standard, m)
	// minStart, maxEnd := 0, 0
	// for i := 0; i < m; i++ {
	// 	line := strings.Fields(readline(r))
	// 	standards[i].start, _ = strconv.Atoi(line[0])
	// 	standards[i].end, _ = strconv.Atoi(line[1])
	// 	standards[i].value = line[2]
	// }
	// fmt.Println("after read input", time.Since(startTime))

	trie := buildTrie(genes, healthMap)
	// trie := buildTrie(genes, health)
	// fmt.Println("after build Trie", time.Since(startTime))

	buildFails(trie)
	// fmt.Println("after build fails", time.Since(startTime))

	// trie.print()
	min, max := 0, 0
	for i := 0; i < m; i++ {
		line := strings.Fields(readline(r))
		var st standard
		st.start, _ = strconv.Atoi(line[0])
		st.end, _ = strconv.Atoi(line[1])
		st.value = line[2]

		total := trie.walkWith(st.value, i, st.start, st.end)
		if i == 0 {
			min, max = total, total
		} else {
			if total > max {
				max = total
			}
			if total < min {
				min = total
			}
		}
	}
	fmt.Printf("%d %d\n", min, max)
	fmt.Println(time.Since(startTime))
	// pprof.WriteHeapProfile(os.Stdout)
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

type healthInfo struct {
	parent *healthInfo
	hVal   [][2]int
	total  map[int]int
}

type node struct {
	char     byte
	children [26]*node
	parent   *node
	fail     *node
	end      bool
	health   *healthInfo
}

func newTrie() *node {
	root := new(node)
	root.fail = root
	return root
}

func (n *node) setEnd(isEnd bool) {
	n.end = isEnd
}

func newHealthInfo(size int) *healthInfo {
	// hVal := make([]int, size)
	total := make(map[int]int)
	return &healthInfo{parent: nil, total: total}
}

func (n *node) setHealth(idx, size int, health [][2]int) {
	if n.health == nil {
		n.health = newHealthInfo(size)
	}
	n.health.hVal = health
}

// func (n *node) setHealth(idx, health, size int) {
// 	if n.health == nil {
// 		n.health = newHealthInfo(size)
// 	}
// 	n.health.hVal[idx] = health
// }

func (n *node) addChild(c byte) *node {
	child := n.children[indexOf(c)]
	if child == nil {
		child = new(node)
		child.char = c
		child.parent = n
		n.children[indexOf(c)] = child
	}
	return child
}

func (n *node) print() {
	q := make([]*node, 0)
	q = append(q, n)
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		fmt.Printf("%c fails to: %+v  ", node.char, node.fail)
		fmt.Printf("%v ", node.health)

		for _, child := range node.children {
			if child != nil {
				q = append(q, child)
			}
		}
		fmt.Println()
	}
}

func indexOf(c byte) int {
	return int(c - 'a')
}

func buildTrie(genes []string, healthMap map[string][][2]int) *node {
	root := newTrie()
	for i, g := range genes {
		n := root
		for j := 0; j < len(g); j++ {
			n = n.addChild(g[j])
			if j == len(g)-1 {
				n.setEnd(true)
				n.setHealth(i, len(genes), healthMap[genes[i]])
			}
		}
	}
	return root
}

func buildFails(root *node) {
	q := make([]*node, 0)
	q = append(q, root)
	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		for _, child := range n.children {
			if child != nil {
				q = append(q, child)
			}
		}
		if n == root {
			continue
		}
		if n.parent == root {
			n.fail = root
		} else {
			f := n.parent.fail
			for {
				if tmp := f.children[indexOf(n.char)]; tmp != nil {
					n.fail = tmp
					if tmp.health != nil {
						if n.health != nil {
							n.health.parent = tmp.health
						} else {
							n.health = tmp.health
						}
					}
					break
				} else if f == root {
					n.fail = root
					break
				}
				f = f.fail
			}
		}
	}
}

func genSum(counts [][2]int, lo, hi int) int {
	sum := 0
	for _, c := range counts {
		if c[0] <= hi && c[0] >= lo {
			sum += c[1]
		}
	}
	// if lo == 0 {
	// 	sum = counts[hi]
	// } else {
	// 	sum = counts[hi] - counts[lo-1]
	// }
	return sum
}

func (h *healthInfo) getSum(stIdx, lo, hi int) int {
	if v, ok := h.total[stIdx]; ok {
		return v
	}
	if h.parent == nil {
		h.total[stIdx] = genSum(h.hVal, lo, hi)
		return h.total[stIdx]
	}
	sum := h.parent.getSum(stIdx, lo, hi)
	sum += genSum(h.hVal, lo, hi)
	h.total[stIdx] = sum
	return sum
}

func (n *node) walkWith(text string, stIdx, lo, hi int) int {
	total := 0
	for i := 0; i < len(text); i++ {
		c := text[i]
		if child := n.children[indexOf(c)]; child != nil {
			if child.health != nil {
				total += child.health.getSum(stIdx, lo, hi)
			}
			n = child
		} else {
			f := n.fail
			for {
				if tmp := f.children[indexOf(c)]; tmp != nil {
					n = tmp
					if tmp.health != nil {
						total += tmp.health.getSum(stIdx, lo, hi)
					}
					break
				} else if f.parent == nil {
					n = f
					break
				}
				f = f.fail
			}
		}
	}
	// fmt.Println(lo, hi, total)
	return total
}
