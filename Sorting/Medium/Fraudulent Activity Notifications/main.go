package main

import "fmt"

func main() {
	var n, d int32
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &d)
	// fmt.Scanln()
	inputs := make([]int32, n)
	for i := 0; i < int(n); i++ {
		var num int32
		fmt.Scanf("%d", &num)
		inputs[i] = num
	}
	// fmt.Println(inputs)
	fmt.Println(activityNotifications(inputs, d))
}

func activityNotifications(expenditure []int32, d int32) int32 {
	cnt := int32(0)
	counts := countSort(expenditure[:d])
	// fmt.Println(expenditure[:d], expenditure[d])
	// fmt.Println(counts)
	m1, m2 := median(expenditure[:d], counts)
	// fmt.Println(m1, m2)
	if expenditure[d] >= m1+m2 {
		cnt++
	}
	for i := d; i < int32(len(expenditure))-1; i++ {
		updateCounts(counts, expenditure[i-d], expenditure[i])
		// fmt.Println(expenditure[i-d+1:i+1], expenditure[i+1])
		// fmt.Println(counts)
		// [Min] 替换的元素和删除的元素在中位数的同侧，则不影响中位数的值，无需更新
		if !(expenditure[i-d] < m1 && expenditure[i] < m1 ||
			expenditure[i-d] > m2 && expenditure[i] > m2) {
			m1, m2 = median(expenditure[i-d+1:i+1], counts)
			// fmt.Println(m1, m2)
		}
		if expenditure[i+1] >= m1+m2 {
			cnt++
		}
	}
	return cnt
}

// 计算 a 的中位数无需对 a 进行实际排序，只需要知道 a 中各元素的 count 即可
func median(a, counts []int32) (m1, m2 int32) {
	l := int32(len(a))
	if l == 0 {
		return 0, 0
	}
	start := int32(0)
	for i := 0; i < len(counts); i++ {
		if counts[i]+start > l/2 && start <= l/2 {
			m2 = int32(i)
			break
		}
		start += counts[i]
	}
	if l%2 == 1 {
		m1 = m2
	} else {
		start = int32(0)
		for i := 0; i < len(counts); i++ {
			if counts[i]+start > l/2-1 && start <= l/2-1 {
				m1 = int32(i)
				break
			}
			start += counts[i]
		}
	}
	return
}

func countSort(a []int32) []int32 {
	counts := make([]int32, 201)
	for _, v := range a {
		counts[v]++
	}
	// [Min] 无需实际排序，只要获取元素值对应的个数即可满足求得中位数的条件
	// for i := 1; i < len(counts); i++ {
	// 	counts[i] += counts[i-1]
	// }
	// for i := len(a) - 1; i >= 0; i-- {
	// 	b[counts[a[i]]-1] = a[i]
	// 	counts[a[i]]--
	// }
	return counts
}

// 根据删除的元素值和替换的元素值更新对应的 count
func updateCounts(counts []int32, delKey, addKey int32) {
	done := 0
	for i := range counts {
		if int32(i) == delKey {
			counts[i]--
			done++
		}
		if int32(i) == addKey {
			counts[i]++
			done++
		}
		if done == 2 {
			return
		}
	}
}
