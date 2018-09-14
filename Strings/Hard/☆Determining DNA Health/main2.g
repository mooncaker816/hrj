package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var primeRK = 23

// type gene struct {
// 	g   string
// 	h   int
// 	idx int
// }

type rollingHash struct {
	hash int
	// next int
	pow  int
	size int
	// health int
}

type standard struct {
	start, end int
	value      string
}

func main() {
	// startTime := time.Now()
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	genes := strings.Fields(readline(r))
	health := strings.Fields(readline(r))
	_ = health
	m, _ := strconv.Atoi(readline(r))
	standards := make([]standard, m)
	// hashMaps := make([]map[int]hashInfo, m)
	hashMap := make(map[int]map[int]int, n) //outside key is hash, inside key is index, value is health for that index
	// hashPool := make(map[int]hashInfo)
	for i := 0; i < m; i++ {
		line := strings.Fields(readline(r))
		standards[i].start, _ = strconv.Atoi(line[0])
		standards[i].end, _ = strconv.Atoi(line[1])
		standards[i].value = line[2]
	}

	// build gene hashes -> hashMap
	seen := make(map[string]int)
	sizeMap := make(map[int]struct{})
	for i, g := range genes {
		sizeMap[len(genes[i])] = struct{}{}
		h, _ := strconv.Atoi(health[i])
		if v, ok := seen[g]; ok {
			m := hashMap[v]
			m[i] = h
			continue
		}
		hash := genHash(genes[i])
		seen[g] = hash
		m := make(map[int]int)
		m[i] = h
		hashMap[hash] = m
	}
	sort.Slice(standards, func(i, j int) bool {

		return standards[i].start < standards[j].start
	})
	// fmt.Println(sizeMap)
	// fmt.Println(time.Since(startTime))
	// fmt.Println(hashMap)
	//
	min, max := 0, 0
	for idx, st := range standards {
		// fmt.Println(idx, time.Since(startTime))
		out := make(chan int)
		var wg sync.WaitGroup
		wg.Add(len(sizeMap))
		for size := range sizeMap {
			size := size
			go func() {
				total := 0
				var r rollingHash
				r.hash = genHash(st.value[:size])
				// r.next = size
				r.size = size
				r.pow = genPow(size)
				// match
				if m, ok := hashMap[r.hash]; ok {
					// _ = m
					for i, h := range m {
						if i >= st.start && i <= st.end {
							total += h
						}
					}
					// total += computeHealth(health, v.indexes)
				}
				for j := size; j < len(st.value); j++ {
					r.hash = r.hash*primeRK + int(st.value[j])
					r.hash -= r.pow * int(st.value[j-r.size])
					if m, ok := hashMap[r.hash]; ok {
						for i, h := range m {
							if i >= st.start && i <= st.end {
								total += h
							}
						}
					}
				}
				out <- total
				wg.Done()
			}()
		}
		go func() {
			wg.Wait()
			close(out)
		}()
		total := 0
		for v := range out {
			total += v
		}
		if idx == 0 {
			min, max = total, total
		} else {
			if total < min {
				min = total
			}
			if total > max {
				max = total
			}
		}
	}
	fmt.Printf("%d %d\n", min, max)
	// fmt.Println(time.Since(startTime))
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

func genHash(g string) int {
	h := 0
	for i := 0; i < len(g); i++ {
		h = h*primeRK + int(g[i])
	}
	return h
}

func genPow(size int) int {
	pow, sq := 1, primeRK
	for i := size; i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return pow
}

func computeHealth(health []string, indexes []int) int {
	var ret int
	for _, i := range indexes {
		hStr := health[i]
		num, _ := strconv.Atoi(hStr)
		ret += num
	}
	return ret
}
