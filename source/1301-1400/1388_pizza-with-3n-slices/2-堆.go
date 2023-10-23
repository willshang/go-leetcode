package main

import "container/heap"

func main() {

}

var m map[int][2]int

func maxSizeSlices(slices []int) int {
	n := len(slices)
	target := n / 3
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	m = make(map[int][2]int)
	for i := 0; i < n; i++ {
		left := (i - 1 + n) % n
		right := (i + 1) % n
		m[i] = [2]int{left, right} // 第i个数左右两边位置
		heap.Push(&intHeap, [2]int{slices[i], i})
	}
	res := 0
	visited := make([]bool, n)
	for i := 0; i < target; {
		top := heap.Pop(&intHeap).([2]int)
		value, index := top[0], top[1]
		if visited[index] == true { // 当前序号不可用
			continue
		}
		i++
		left, right := m[index][0], m[index][1]
		visited[left], visited[right] = true, true
		res = res + value
		slices[index] = slices[left] + slices[right] - value // 更新当前序号的值为反悔值
		heap.Push(&intHeap, [2]int{slices[index], index})    // 重新赋值放回堆
		reconnect(left)                                      // 更改左边数的指针
		reconnect(right)                                     // 更改右边数的指针
	}
	return res
}

func reconnect(index int) {
	left, right := m[index][0], m[index][1]
	m[right] = [2]int{left, m[right][1]}
	m[left] = [2]int{m[left][0], right}
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][0] > h[j][0]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
