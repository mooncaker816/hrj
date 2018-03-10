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

// old version - timeout in hackerrank but ok in local
// xxxxx

// import (
// 	"fmt"
// 	"math/bits"
// 	"sort"
// )

// type myslice []myint

// type myint struct {
// 	value uint
// 	//len   uint // 从右往左最后一个有效0的位置
// }

// func (s myslice) less(i, j int) bool {
// 	if s[i].value < s[j].value {
// 		return true
// 	}
// 	return false
// }

// func main() {
// 	var t int
// 	fmt.Scanf("%d\n", &t)
// 	s0 := make(myslice, 0, t)
// 	s := make(myslice, 0, t)
// 	var tmplen int
// 	for i := 0; i < t; i++ {
// 		var n uint
// 		fmt.Scanf("%d", &n)
// 		if i == 0 {
// 			tmplen = bits.Len(n)
// 			//s = append(s, myint{n})
// 		}
// 		if bits.Len(n) == tmplen {
// 			s = append(s, myint{n})
// 		} else if bits.Len(n) < tmplen {
// 			s0 = append(s0, myint{n})
// 		} else if bits.Len(n) > tmplen {
// 			s0 = append(s0, s...)
// 			s = s[0:0]
// 			s = append(s, myint{n})
// 			tmplen = bits.Len(n)
// 		}

// 		//prefix := bits.LeadingZeros(n)
// 		//s[i].pos = uint(bits.Len(uint((^uint(0) - bits.RotateLeft(s[i].value, prefix)) >> uint(prefix))))
// 		//s[i].len = uint(bits.Len(n))
// 	}
// 	min := ^uint(0)
// 	if len(s0) > 0 {
// 		for _, v0 := range s0 {
// 			for _, v := range s {
// 				if xord := v0.value ^ v.value; xord < min {
// 					min = xord
// 				}
// 			}
// 		}
// 		fmt.Println(min)
// 		return
// 	}
// 	sort.Slice(s, s.less)
// 	a := s[0].value
// 	b := s[t-1].value
// 	maxlen := bits.Len(b)
// 	common := 0
// 	for {
// 		if bits.Len(a) < bits.Len(b) {
// 			break
// 		}
// 		a -= 1 << uint(bits.Len(a)-1)
// 		b -= 1 << uint(bits.Len(b)-1)
// 		common++
// 	}
// 	//var prev uint
// 	//fmt.Println(common)

// 	//prev = s[0].len
// 	next := -1
// 	for i := 0; i < t; i++ {
// 		if next != -1 && i >= next {
// 			break
// 		}
// 		for j := t - 1; j >= 0; j-- {
// 			if next == -1 {
// 				if bits.OnesCount((s[j].value >> uint(maxlen-common-1))) != common+1 {
// 					next = j + 1
// 					break
// 				}
// 			} else {
// 				if j < next {
// 					break
// 				}
// 			}
// 			xord := s[i].value ^ s[j].value
// 			if xord < min {
// 				min = xord
// 			}
// 		}
// 		//prev = s[i].len
// 	}

// 	//fmt.Println(next)
// 	// for i := 0; i < next; i++ {
// 	// 	for j := next; j < t; j++ {
// 	// 		xord := s[i].value ^ s[j].value
// 	// 		if xord < min {
// 	// 			min = xord
// 	// 		}
// 	// 	}
// 	// }
// 	fmt.Println(min)
// }
