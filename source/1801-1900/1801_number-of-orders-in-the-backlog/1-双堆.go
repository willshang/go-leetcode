package main

import "container/heap"

func main() {

}

// leetcode1801_积压订单中的订单总数
func getNumberOfBacklogOrders(orders [][]int) int {
	res := 0
	buyHeap := make(BuyHeap, 0)
	sellHeap := make(SellHeap, 0)
	heap.Init(&buyHeap)
	heap.Init(&sellHeap)
	for i := 0; i < len(orders); i++ {
		price := orders[i][0]
		count := orders[i][1]
		typeInt := orders[i][2]
		if typeInt == 1 { // sell
			for buyHeap.Len() > 0 {
				node := heap.Pop(&buyHeap).(Node)
				if node.price < price {
					heap.Push(&buyHeap, node)
					break
				}
				if node.count > count { // 数量大于
					node.count = node.count - count
					count = 0
					heap.Push(&buyHeap, node)
					break
				}
				count = count - node.count
			}
			if count > 0 {
				heap.Push(&sellHeap, Node{
					count: count,
					price: price,
				})
			}
		} else { // buy
			for sellHeap.Len() > 0 {
				node := heap.Pop(&sellHeap).(Node)
				if node.price > price {
					heap.Push(&sellHeap, node)
					break
				}
				if node.count > count { // 数量小于
					node.count = node.count - count
					count = 0
					heap.Push(&sellHeap, node)
					break
				}
				count = count - node.count
			}
			if count > 0 {
				heap.Push(&buyHeap, Node{
					count: count,
					price: price,
				})
			}
		}
	}
	for buyHeap.Len() > 0 {
		node := heap.Pop(&buyHeap).(Node)
		res = res + node.count
	}
	for sellHeap.Len() > 0 {
		node := heap.Pop(&sellHeap).(Node)
		res = res + node.count
	}
	return res % 1000000007
}

type Node struct {
	count int
	price int
}

type SellHeap []Node

func (h SellHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h SellHeap) Less(i, j int) bool {
	return h[i].price < h[j].price
}

func (h SellHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *SellHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *SellHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

type BuyHeap []Node

func (h BuyHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h BuyHeap) Less(i, j int) bool {
	return h[i].price > h[j].price
}

func (h BuyHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *BuyHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *BuyHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
