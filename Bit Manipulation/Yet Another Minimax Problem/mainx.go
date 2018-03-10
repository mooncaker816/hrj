 xxxxx

import (
	"fmt"
	"math/bits"
	"sort"
)

type myslice []myint

type myint struct {
	value uint
	//len   uint // 从右往左最后一个有效0的位置
}

func (s myslice) less(i, j int) bool {
	if s[i].value < s[j].value {
		return true
	}
	return false
}

func main() {
	var t int
	fmt.Scanf("%d\n", &t)
	s0 := make(myslice, 0, t)
	s := make(myslice, 0, t)
	var tmplen int
	for i := 0; i < t; i++ {
		var n uint
		fmt.Scanf("%d", &n)
		if i == 0 {
			tmplen = bits.Len(n)
			//s = append(s, myint{n})
		}
		if bits.Len(n) == tmplen {
			s = append(s, myint{n})
		} else if bits.Len(n) < tmplen {
			s0 = append(s0, myint{n})
		} else if bits.Len(n) > tmplen {
			s0 = append(s0, s...)
			s = s[0:0]
			s = append(s, myint{n})
			tmplen = bits.Len(n)
		}

		//prefix := bits.LeadingZeros(n)
		//s[i].pos = uint(bits.Len(uint((^uint(0) - bits.RotateLeft(s[i].value, prefix)) >> uint(prefix))))
		//s[i].len = uint(bits.Len(n))
	}
	min := ^uint(0)
	if len(s0) > 0 {
		for _, v0 := range s0 {
			for _, v := range s {
				if xord := v0.value ^ v.value; xord < min {
					min = xord
				}
			}
		}
		fmt.Println(min)
		return
	}
	sort.Slice(s, s.less)
	a := s[0].value
	b := s[t-1].value
	maxlen := bits.Len(b)
	common := 0
	for {
		if bits.Len(a) < bits.Len(b) {
			break
		}
		a -= 1 << uint(bits.Len(a)-1)
		b -= 1 << uint(bits.Len(b)-1)
		common++
	}
	//var prev uint
	//fmt.Println(common)

	//prev = s[0].len
	next := -1
	for i := 0; i < t; i++ {
		if next != -1 && i >= next {
			break
		}
		for j := t - 1; j >= 0; j-- {
			if next == -1 {
				if bits.OnesCount((s[j].value >> uint(maxlen-common-1))) != common+1 {
					next = j + 1
					break
				}
			} else {
				if j < next {
					break
				}
			}
			xord := s[i].value ^ s[j].value
			if xord < min {
				min = xord
			}
		}
		//prev = s[i].len
	}

	//fmt.Println(next)
	// for i := 0; i < next; i++ {
	// 	for j := next; j < t; j++ {
	// 		xord := s[i].value ^ s[j].value
	// 		if xord < min {
	// 			min = xord
	// 		}
	// 	}
	// }
	fmt.Println(min)
}
