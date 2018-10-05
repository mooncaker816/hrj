package main

import (
	"container/heap"
)

type item struct {
	index int
	val   int
	high  bool
}

type priorityQueue []*item

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].val < pq[j].val }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *priorityQueue) update(item *item, val int) {
	item.val = val
	heap.Fix(pq, item.index)
}

type pql struct {
	priorityQueue
}

type pqh struct {
	priorityQueue
}

func (pq pqh) Less(i, j int) bool { return pq.priorityQueue[i].val > pq.priorityQueue[j].val }

func activityNotifications2(expenditure []int, d int) int {
	odd := d&1 == 1
	orig, pql, pqh := initPush(expenditure[:d])
	m := findMedian(pql, pqh, odd)
	cnt := 0
	for i := d; i < len(expenditure); i++ {
		if expenditure[i] >= m {
			cnt++
		}
		remove(pql, pqh, orig[i%d])
		new := new(item)
		new.val = expenditure[i]
		orig[i%d] = new
		add(pql, pqh, new)
		m = findMedian(pql, pqh, odd)
	}
	return cnt
}

func initPush(a []int) ([]*item, *pql, *pqh) {
	orig := make([]*item, len(a))
	pql := new(pql)
	pqh := new(pqh)
	pql.priorityQueue = make(priorityQueue, 0, len(a))
	pqh.priorityQueue = make(priorityQueue, 0, len(a))

	for i, v := range a {
		item := &item{val: v}
		orig[i] = item
		add(pql, pqh, item)
	}
	return orig, pql, pqh
}

func add(pql *pql, pqh *pqh, item *item) {
	if pqh.Len() == 0 || item.val <= pqh.priorityQueue[0].val {
		// if pqh.Len() == 0 {
		item.high = true
		heap.Push(pqh, item)
		rebalance(pql, pqh)
		return
	}
	item.high = false
	heap.Push(pql, item)
	rebalance(pql, pqh)
}

func remove(pql *pql, pqh *pqh, item *item) {
	if item.high {
		heap.Remove(pqh, item.index)
	} else {
		heap.Remove(pql, item.index)
	}
	rebalance(pql, pqh)
}

func rebalance(pql *pql, pqh *pqh) {
	hl, ll := pqh.Len(), pql.Len()
	if hl > ll+1 {
		item := heap.Pop(pqh).(*item)
		item.high = false
		heap.Push(pql, item)
	} else if ll > hl+1 {
		item := heap.Pop(pql).(*item)
		item.high = true
		heap.Push(pqh, item)
	}
}

func findMedian(pql *pql, pqh *pqh, odd bool) (m int) {
	hl, ll := pqh.Len(), pql.Len()
	if hl == ll {
		return pqh.priorityQueue[0].val + pql.priorityQueue[0].val
	} else if hl > ll {
		return 2 * pqh.priorityQueue[0].val
	}
	return 2 * pql.priorityQueue[0].val
}
