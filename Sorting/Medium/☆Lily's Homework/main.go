package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int32
	fmt.Scanf("%d\n", &n)
	inputs := make([]int32, n)
	for i := 0; i < int(n); i++ {
		var num int32
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}
	fmt.Println(lilysHomework(inputs))
}

func lilysHomework(arr []int32) int32 {
	cnt1, cnt2 := int32(0), int32(0)
	arrM1, arrM2, arrCopy1, arrCopy2, arr2 := preload(arr)
	// 升序排列
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	for i, v := range arrCopy1 {
		// 与目标值不同，需要交换
		if v != arr[i] {
			cnt1++
			// 获取目标值在当前未排序的切片中的索引 j
			j := arrM1[arr[i]]
			// 交换 i,j 两个元素，此时 i 元素已经归位，j 上的元素为交换之前的 i 的值，
			// 等待后续比较 j 的时候判断是否需要继续交换
			arrCopy1[i], arrCopy1[j] = arrCopy1[j], arrCopy1[i]
			arrM1[v] = j
		}
	}
	// 降序
	sort.Slice(arr2, func(i, j int) bool {
		return arr2[i] > arr2[j]
	})
	for i, v := range arrCopy2 {
		// 与目标值不同，需要交换
		if v != arr2[i] {
			cnt2++
			// 获取目标值在当前未排序的切片中的索引 j
			j := arrM2[arr2[i]]
			// 交换 i,j 两个元素，此时 i 元素已经归位，j 上的元素为交换之前的 i 的值，
			// 等待后续比较 j 的时候判断是否需要继续交换
			arrCopy2[i], arrCopy2[j] = arrCopy2[j], arrCopy2[i]
			arrM2[v] = j
		}
	}
	if cnt1 < cnt2 {
		return cnt1
	}
	return cnt2
}

func preload(arr []int32) (map[int32]int32, map[int32]int32, []int32, []int32, []int32) {
	m1 := make(map[int32]int32)
	m2 := make(map[int32]int32)
	copy1 := make([]int32, len(arr))
	copy2 := make([]int32, len(arr))
	copy3 := make([]int32, len(arr))
	for i, v := range arr {
		m1[v] = int32(i)
		m2[v] = int32(i)
		copy1[i] = v
		copy2[i] = v
		copy3[i] = v
	}
	return m1, m2, copy1, copy2, copy3
}
