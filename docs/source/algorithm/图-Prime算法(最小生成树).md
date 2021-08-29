# 图-Prime算法

## 1、定义

- 最小生成树-Prime（普里姆算法）
- 初始输入：一个加权连通图，顶点集合：V，边集合：E；都为空
- 1、任取一个顶点加入生成树，即队列V
- 2、在那些一个端点在生成树里，另一个端点不在生成树里的边中，选取一条权最小的边，将它和另一个端点加进生成树
- 3、重复2
- 适合：稠密图
- 时间复杂度：

## 2、Go实现

- Prime

```go
package main

import "math"

func main() {

}

// Prime：传入邻接矩阵
func Prime(arr [][]int) int {
	res := 0
	n := len(arr)
	visited := make([]bool, n)
	target := 0
	visited[target] = true // 任选1个即可
	dis := make([]int, n)  // 任意选择的节点：到其它点的距离
	for i := 0; i < n; i++ {
		dis[i] = arr[target][i]
	}
	for i := 1; i < n; i++ { // 执行n-1次：n-1条边
		var index int
		minValue := math.MaxInt32
		for j := 0; j < n; j++ { // 寻找：未访问过的最短边
			if visited[j] == false && dis[j] < minValue {
				minValue = dis[j]
				index = j
			}
		}
		visited[index] = true    // 标记为访问过的点
		res = res + minValue     // 加上最短边
		for j := 0; j < n; j++ { // 更新距离：以index为起点，更新生成树到每一个非树顶点的距离
			if visited[j] == false && dis[j] > arr[index][j] {
				dis[j] = arr[index][j]
			}
		}
	}
	return res
}
```

- Prime堆优化

```go
package main

import (
	"container/heap"
	"math"
)

func main() {

}

// Prime-堆优化-邻接表
func Prime(arr [][]int) int {
	res := 0
	n := len(arr)
	visited := make([]bool, n)
	target := 0
	dis := make([]int, n) // 任意选择的节点：到其它点的距离
	for i := 0; i < n; i++ {
		dis[i] = math.MaxInt32
	}
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{0, target}) // [2]int{距离，下标}
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([2]int)
		minValue, index := node[0], node[1]
		if visited[index] == true {
			continue
		}
		visited[index] = true
		res = res + minValue
		for j := 0; j < len(arr[index]); j++ {
			if visited[j] == false && dis[j] > arr[index][j] {
				dis[j] = arr[index][j]
				heap.Push(&intHeap, [2]int{arr[index][j], j})
			}
		}
	}
	return res
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][0] < h[j][0]
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
```

## 3、Leetcode

| Title                                                        | Tag                      | 难度   | 完成情况 |
| ------------------------------------------------------------ | ------------------------ | ------ | -------- |
| [1584.连接所有点的最小费用](https://leetcode-cn.com/problems/min-cost-to-connect-all-points/) | 并查集、数组、最小生成树 | Medium | 完成     |

