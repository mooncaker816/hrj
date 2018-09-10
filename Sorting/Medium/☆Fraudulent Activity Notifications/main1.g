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
	sortedExp, releaseSeq := countSort(expenditure[:d])
	// fmt.Println("initial: ", sortedExp, releaseSeq)
	m1, m2 := median(sortedExp)
	if expenditure[d] >= m1+m2 {
		cnt++
	}
	for i := d; i < int32(len(expenditure))-1; i++ {
		releaseIdx := releaseSeq[0]
		delta, r0, r1 := int32(0), int32(0), int32(0)
		if expenditure[i] == sortedExp[releaseIdx] {
			sortedExp[releaseIdx] = expenditure[i]
			// fmt.Println("after swap:", sortedExp, releaseSeq)
			// releaseSeq = releaseSeq[1:]
			// releaseSeq = append(releaseSeq, releaseIdx)
		} else if expenditure[i] > sortedExp[releaseIdx] {
			sortedExp[releaseIdx] = expenditure[i]
			if releaseIdx < int32(len(sortedExp)-1) {
				r0 = sortedExp[releaseIdx+1]
				// fmt.Println("after swap:", sortedExp, releaseSeq)
				for j := releaseIdx; j < int32(len(sortedExp))-1 && sortedExp[j] > sortedExp[j+1]; j++ {
					r1 = sortedExp[j+1]
					sortedExp[j], sortedExp[j+1] = sortedExp[j+1], sortedExp[j]
					delta = 1
					// releaseSeq[j+1]--
				}
			}
			// releaseSeq = releaseSeq[1:]
			// releaseSeq = append(releaseSeq, releaseIdx)
		} else {
			sortedExp[releaseIdx] = expenditure[i]
			if releaseIdx >= 1 {
				r1 = sortedExp[releaseIdx-1]
				// fmt.Println("after swap:", sortedExp, releaseSeq)
				for j := releaseIdx; j > 0 && sortedExp[j] < sortedExp[j-1]; j-- {
					r0 = sortedExp[j-1]
					sortedExp[j], sortedExp[j-1] = sortedExp[j-1], sortedExp[j]
					delta = -1
					// releaseSeq[j-1]++
				}
			}
			// releaseSeq = releaseSeq[1:]
			// releaseSeq = append(releaseSeq, releaseIdx)
		}
		if delta != 0 {
			for i := 0; i < len(releaseSeq); i++ {
				if releaseSeq[i] < 0 || releaseSeq[i] >= int32(len(sortedExp)) {
					fmt.Println(r0, r1)
					fmt.Println(releaseIdx)
					fmt.Println("ooooooo:", i, releaseSeq[i], releaseSeq)
				}
				if sortedExp[releaseSeq[i]] <= r1 && sortedExp[releaseSeq[i]] >= r0 {
					releaseSeq[i] -= delta
					if releaseSeq[i] < 0 || releaseSeq[i] >= int32(len(sortedExp)) {
						releaseSeq[i] += delta
					}
				}
			}
		}
		releaseIdx += delta
		copy(releaseSeq[:d-1], releaseSeq[1:d])
		releaseSeq[d-1] = releaseIdx
		// releaseSeq = append(releaseSeq, releaseIdx)

		// fmt.Println("after resort:", sortedExp, releaseSeq)
		m1, m2 = median(sortedExp)
		// fmt.Println(m1, m2, expenditure[i+1])
		if expenditure[i+1] >= m1+m2 {
			cnt++
		}
	}
	return cnt
}

func median(a []int32) (int32, int32) {
	l := len(a)
	if l == 0 {
		return 0, 0
	}
	if l%2 == 1 {
		return a[l/2], a[l/2]
	}
	return a[l/2-1], a[l/2]
}

func countSort(a []int32) ([]int32, []int32) {
	counts := make([]int32, 201)
	b := make([]int32, len(a))
	for _, v := range a {
		counts[v]++
	}
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}
	idxSeq := make([]int32, len(a))
	for i := len(a) - 1; i >= 0; i-- {
		b[counts[a[i]]-1] = a[i]
		counts[a[i]]--
		idxSeq[i] = counts[a[i]]
	}
	// fmt.Println(b, idxSeq)
	return b, idxSeq
}
