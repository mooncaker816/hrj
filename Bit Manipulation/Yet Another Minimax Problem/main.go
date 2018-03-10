package main

import (
	"fmt"
	"math/bits"
	"sync"
)

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	var array [31][]uint
	maxlen := 0
	difflen := false
	prev := 0
	for i := 0; i < t; i++ {
		var n uint
		fmt.Scanf("%d", &n)
		l := bits.Len(n)
		if i == 0 {
			prev = l
		}
		if l > maxlen {
			maxlen = l
		}
		if l != prev {
			difflen = true
		}
		array[l] = append(array[l], n)
	}
	res := ^uint(0)
	if difflen == true {
		ch := make(chan uint)
		done := make(chan struct{})
		var wg sync.WaitGroup
		go func() {
			minall := ^uint(0)
			for v := range ch {
				//fmt.Println("v ", v)
				if v < minall {
					minall = v
				}
			}
			fmt.Println(minall)
			close(done)
		}()
		for i := 0; i < maxlen; i++ {
			if len(array[i]) > 0 {
				// fmt.Println("xxx")
				wg.Add(1)
				go func(a []uint) {
					min := ^uint(0)
					for p := 0; p < len(a); p++ {
						for q := 0; q < len(array[maxlen]); q++ {
							if a[p]^array[maxlen][q] < min {
								min = a[p] ^ array[maxlen][q]
							}
						}
					}
					ch <- min
					wg.Done()
					//fmt.Println("done", min)
				}(array[i])
			}
		}
		wg.Wait()
		close(ch)
		<-done
		return
	}
	tmpl := maxlen
	clearbit := uint(0)
	s0 := make([]uint, 0)
	s1 := make([]uint, 0)
	for tmpl > 0 {
		tmpl--
		clearbit |= 1 << uint(tmpl)
		for i := 0; i < len(array[maxlen]); i++ {
			if bits.Len(array[maxlen][i]&^clearbit) != tmpl {
				s0 = append(s0, array[maxlen][i])
			} else {
				s1 = append(s1, array[maxlen][i])
			}
		}
		if len(s0) > 0 {
			break
		} else {
			s1 = s1[:0]
		}
	}

	for i := 0; i < len(s0); i++ {
		for j := 0; j < len(s1); j++ {
			if s0[i]^s1[j] < res {
				res = s0[i] ^ s1[j]
			}
		}
	}
	if maxlen == 0 {
		res = 0
	}
	fmt.Println(res)
}
