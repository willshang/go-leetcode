package main

import (
	"container/heap"
	"fmt"
)

func main() {
	//fmt.Println(busiestServers(3, []int{1, 2, 3, 4, 5}, []int{5, 2, 3, 3, 3}))
	fmt.Println(busiestServers(3, []int{1, 2, 3, 4}, []int{1, 2, 1, 2}))
}

func busiestServers(k int, arrival []int, load []int) []int {
	n := len(arrival)
	res := make([]int, 0)
	freeHeap := &mixHeap{isBig: false} // 小根堆：下标小的优先处理
	busyHeap := &mixHeap{isBig: false} // 小根堆： 空闲时间小，早结束
	for i := 0; i < k; i++ {
		freeHeap.push([]int{i})
	}
	arr := make([]int, k)
	for i := 0; i < n; i++ {
		start := arrival[i]
		// busy堆执行完出队列
		for busyHeap.Len() > 0 && busyHeap.Top()[0] <= start {
			id := busyHeap.pop()[1]
			// i+((id-i)%k+k)%k 表示下一个循环的编号
			freeHeap.push([]int{i + ((id-i)%k+k)%k}) // 入free堆：注意时间处理，负数取模后要取正
		}
		// 选择一个执行，如果为空舍弃：条件2
		if freeHeap.Len() > 0 {
			id := freeHeap.pop()[0] % k
			busyHeap.push([]int{start + load[i], id}) // 结束时间+id
			arr[id]++
		}
	}
	var maxValue int
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxValue {
			maxValue = arr[i]
			res = []int{i}
		} else if arr[i] == maxValue {
			res = append(res, i)
		}
	}
	return res
}

type mixHeap struct {
	arr   [][]int
	isBig bool
}

func (m *mixHeap) Len() int {
	return len(m.arr)
}

func (m *mixHeap) Swap(i, j int) {
	m.arr[i], m.arr[j] = m.arr[j], m.arr[i]
}

func (m *mixHeap) Less(i, j int) bool {
	if m.isBig {
		return m.arr[i][0] > m.arr[j][0] // 大根堆
	}
	return m.arr[i][0] < m.arr[j][0] // 小根堆
}

func (m *mixHeap) Push(x interface{}) {
	m.arr = append(m.arr, x.([]int))
}

func (m *mixHeap) Pop() interface{} {
	value := (m.arr)[len(m.arr)-1]
	m.arr = (m.arr)[:len(m.arr)-1]
	return value
}

func (m *mixHeap) push(x []int) {
	heap.Push(m, x)
}

func (m *mixHeap) pop() []int {
	return heap.Pop(m).([]int)
}

func (m *mixHeap) Top() []int {
	if m.Len() > 0 {
		return m.arr[0]
	}
	return nil
}
