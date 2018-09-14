package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

type hashInfo struct {
	indexes []int
	// stidx   []int
}

type standard struct {
	start, end int
	value      string
}

func main() {
	startTime := time.Now()
	r := bufio.NewReader(os.Stdin)
	n, _ := strconv.Atoi(readline(r))
	_ = n
	genes := strings.Fields(readline(r))
	health := strings.Fields(readline(r))
	_ = health
	m, _ := strconv.Atoi(readline(r))
	standards := make([]standard, m)
	hashMaps := make([]map[int]hashInfo, m)
	// hashPool := make(map[int]hashInfo)
	for i := 0; i < m; i++ {
		line := strings.Fields(readline(r))
		standards[i].start, _ = strconv.Atoi(line[0])
		standards[i].end, _ = strconv.Atoi(line[1])
		standards[i].value = line[2]
		hashMaps[i] = make(map[int]hashInfo, 0) //key is hash
	}

	// build gene hashes -> hashMap
	// sort.Slice(standards, func(i, j int) bool {
	// 	return standards[i].start < standards[j].start
	// })
	// start, end := 0, 0
	seen := make(map[string]int)
	sizeMap := make(map[int]struct{})
	for idx, st := range standards {
		// if st.start > end {
		// 	start, end = st.start, st.end
		// } else if st.start == end {
		// 	start, end = st.start+1, st.end
		// } else {
		// 	if st.end <= end {
		// 		continue
		// 	}
		// 	start, end = end+1, st.end
		// }
		for i := st.start; i <= st.end; i++ {
			sizeMap[len(genes[i])] = struct{}{}
			if v, ok := seen[genes[i]]; ok {
				hi := hashMaps[idx][v]
				// hi := hashPool[v]
				hi.indexes = append(hi.indexes, i)
				hashMaps[idx][v] = hi
				continue
			}
			indexes := make([]int, 1)
			indexes[0] = i
			h := genHash(genes[i])
			seen[genes[i]] = h
			// hashPool[h] = hashInfo{indexes: indexes}
			hashMaps[idx][h] = hashInfo{indexes: indexes}
			// hashMap[h].stidx = stidx
		}
	}
	fmt.Println(sizeMap)
	fmt.Println(time.Since(startTime))
	//
	// min, max := 0, 0
	// for idx, st := range standards {
	// 	out := make(chan int)
	// 	var wg sync.WaitGroup
	// 	wg.Add(len(sizeMap))
	// 	for size := range sizeMap {
	// 		size := size
	// 		go func() {
	// 			total := 0
	// 			var r rollingHash
	// 			r.hash = genHash(st.value[:size])
	// 			// r.next = size
	// 			r.size = size
	// 			r.pow = genPow(size)
	// 			// match
	// 			if v, ok := hashMaps[idx][r.hash]; ok {
	// 				total += computeHealth(health, v.indexes)
	// 			}
	// 			for j := size; j < len(st.value); j++ {
	// 				r.hash = r.hash*primeRK + int(st.value[j])
	// 				r.hash -= r.pow * int(st.value[j-r.size])
	// 				if v, ok := hashMaps[idx][r.hash]; ok {
	// 					total += computeHealth(health, v.indexes)
	// 				}
	// 			}
	// 			out <- total
	// 			wg.Done()
	// 		}()
	// 	}
	// 	go func() {
	// 		wg.Wait()
	// 		close(out)
	// 	}()
	// 	total := 0
	// 	for v := range out {
	// 		total += v
	// 	}
	// 	if idx == 0 {
	// 		min, max = total, total
	// 	} else {
	// 		if total < min {
	// 			min = total
	// 		}
	// 		if total > max {
	// 			max = total
	// 		}
	// 	}
	// }
	// fmt.Printf("%d %d\n", min, max)
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
