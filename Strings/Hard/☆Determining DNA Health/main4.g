package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var primeRK = 23

// case0 : 0 19
// case1 : 15806635 20688978289
// case2 : 0 3120724375
// case3 : 3218660 11137051
// case4 : 207177686 3128254990

type rollingHash struct {
	hash int
	// next int
	pow int
	// size int
	// health int
}

type standard struct {
	start, end int
	value      string
}

type hashInfo struct {
	pattern        string
	counts         []int
	tmpIdx, tmpSum int
}

func main() {
	startTime := time.Now()
	_ = startTime
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	_ = n
	genes := strings.Fields(readline(r))
	health := strings.Fields(readline(r))
	_ = health
	m, _ := strconv.Atoi(readline(r))
	standards := make([]standard, m)
	minStart, maxEnd := 0, 0
	for i := 0; i < m; i++ {
		line := strings.Fields(readline(r))
		standards[i].start, _ = strconv.Atoi(line[0])
		standards[i].end, _ = strconv.Atoi(line[1])
		standards[i].value = line[2]
		if standards[i].start < minStart {
			minStart = standards[i].start
		}
		if standards[i].end > maxEnd {
			maxEnd = standards[i].end
		}
	}
	hashMap := make(map[int]hashInfo, maxEnd-minStart+1) // key is hash,  value is {gene string,health count}

	// fmt.Println("minStart in standards:", minStart)
	// fmt.Println("maxEnd in standards:", maxEnd)
	// fmt.Println("len of genes:", n)
	// fmt.Println("before initial hashMap", time.Since(startTime))

	// build gene hashes -> hashMap
	seen := make(map[string]int)
	sizeMap := make(map[int]struct{})
	for i := minStart; i <= maxEnd; i++ {
		sizeMap[len(genes[i])] = struct{}{}
		h, _ := strconv.Atoi(health[i])
		if v, ok := seen[genes[i]]; ok {
			// if hashMap[v].pattern != genes[i] {
			// 	fmt.Println("collision", hashMap[v].pattern, genes[i])
			// }
			hi := hashMap[v]
			hi.counts[i] = h
			continue
		}
		var hi hashInfo
		hi.pattern = genes[i]
		hash := genHash(genes[i])
		hi.counts = make([]int, maxEnd+1)
		hi.counts[i] = h
		hashMap[hash] = hi
		seen[genes[i]] = hash
	}
	// fmt.Println("after initial hashMap", time.Since(startTime))
	// fmt.Println(hashMap)

	var wg sync.WaitGroup
	wg.Add(len(hashMap))
	for _, hi := range hashMap {
		go func(hi hashInfo) {
			for i := 1; i < len(hi.counts); i++ {
				hi.counts[i] += hi.counts[i-1]
			}
			wg.Done()
		}(hi)
	}
	wg.Wait()

	// fmt.Println("after rebuild counts", time.Since(startTime))
	// fmt.Println(sizeMap)
	// fmt.Println(hashMap)
	//
	min, max := 0, 0

	// calculated := make(map[string]int, len(hashMap))
	for idx, st := range standards {
		total := 0
		for size := range sizeMap {
			var r rollingHash
			// if hash, ok := calculated[st.value[:size]]; ok {
			// 	r.hash = hash
			// } else {
			r.hash = genHash(st.value[:size])
			// calculated[st.value[:size]] = r.hash
			// }
			// r.next = size
			// r.size = size
			r.pow = genPow(size)
			// match
			// if hi, ok := hashMap[r.hash]; ok && st.value[:size] == hi.pattern {
			if hi, ok := hashMap[r.hash]; ok {
				if st.start == 0 {
					total += hi.counts[st.end]
				} else {
					total += hi.counts[st.end] - hi.counts[st.start-1]
				}
			}
			for j := size; j < len(st.value); j++ {
				r.hash = r.hash*primeRK + int(st.value[j])
				r.hash -= r.pow * int(st.value[j-size])
				// if hi, ok := hashMap[r.hash]; ok && st.value[j-size+1:j+1] == hi.pattern {
				if hi, ok := hashMap[r.hash]; ok {
					if st.start == 0 {
						total += hi.counts[st.end]
					} else {
						total += hi.counts[st.end] - hi.counts[st.start-1]
					}
				}
			}
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
