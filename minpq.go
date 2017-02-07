package main

import (
    "container/heap"
)

// An Item is something we manage in a priority queue.
type Node struct {
    label int // The label of the node
    distTo float64    // The priority of the item in the queue.
    // The index is needed by update and is maintained by the heap.Interface methods.
    index int // The index of the item in the heap.
}

// A MinPQ implements heap.Interface and holds Items.
type MinPQ []*Node

func (pq MinPQ) Len() int {
    return len(pq)
}

// We want Pop to give us the shortest path here found
func (pq MinPQ) Less(i, j int) bool {
    return pq[i].distTo < pq[j].distTo
}

func (pq MinPQ) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}

func (pq *MinPQ) Push(x interface{}) {
    n := len(*pq)
    node := x.(*Node)
    node.index = n
    *pq = append(*pq, node)
}

func (pq *MinPQ) Pop() interface{} {
    old := *pq
    n := len(old)
    node := old[n-1]
    node.index = -1 // for safety
    *pq = old[0 : n-1]
    return node
}

// update modifies the priority and value of an Item in the queue.
func (pq *MinPQ) update(node *Node, index int) {
    node.index = index
    heap.Fix(pq, index)
}
