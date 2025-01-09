package linked_list

import (
	"container/heap"
	"leetcode/tasks/custom_structs"
)

/*
Задача:

	Дано k отсортированных списков, необходимо смерджить их и в результате получить отсортированный список
*/
func MergeKLists(lists []*custom_structs.ListNode) *custom_structs.ListNode {
	pq := make(PQ, 0)
	for _, node := range lists {
		if node != nil {
			pq = append(pq, node)
		}
	}

	if len(pq) == 0 {
		return nil
	}
	heap.Init(&pq)

	head := &custom_structs.ListNode{}
	dummyHead := head

	for len(pq) > 0 {
		min := heap.Pop(&pq)
		minNode := min.(*custom_structs.ListNode)
		head.Next = minNode
		head = head.Next

		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}
	return dummyHead.Next
}

type PQ []*custom_structs.ListNode

func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
}

func (pq PQ) Less(a, b int) bool {
	return pq[a].Val < pq[b].Val
}

func (pq *PQ) Push(nodeInterface interface{}) {
	node := nodeInterface.(*custom_structs.ListNode)
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	lastNode := old[len(*pq)-1]
	*pq = old[:len(*pq)-1]
	return lastNode
}

