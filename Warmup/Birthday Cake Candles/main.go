package main

import (
	"fmt"
	"sort"
)

type Int64Slice []int64

func (s Int64Slice) Len() int           { return len(s) }
func (s Int64Slice) Less(i, j int) bool { return s[i] < s[j] }
func (s Int64Slice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	var n int64
	fmt.Scanf("%d\n", &n)
	s := Int64Slice(make([]int64, n))
	for i := 0; int64(i) < n; i++ {
		fmt.Scanf("%d", &s[i])
	}
	//size, _ := strconv.ParseInt(r.Text(), 10, 64)
	// r.Scan()
	// s := strings.Fields(r.Text())
	// var nums Int64Slice
	// for _, v := range s {
	// 	val, _ := strconv.ParseInt(v, 10, 64)
	// 	nums = append(nums, val)
	// }
	var num int64
	sort.Stable(s)
	max := s[n-1]
	for _, v := range s {
		if v == max {
			num++
		}
	}
	fmt.Println(num)
}
