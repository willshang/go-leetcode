# 图-最短路径算法-Dijkstra

## 0、定义

- **解决问题**：从图中的某个顶点出发到达另外一个顶点的所经过的边的权重和最小的一条路径，称为最短路径。
- 算法：迪杰斯特拉算法主要特点是**从起始点开始**，采用**贪心算法的策略**，每次遍历到始点距离最近且未访问过的顶点的邻接节点，**直到扩展到终点为止。**

## 1、操作

## 2、Go实现

- 不加堆实现

```go
// Dijkstra：index => 其它地址距离
// arr：邻接矩阵
func Dijkstra(arr [][]int, index int) []int {
	n := len(arr)
	maxValue := math.MaxInt32
	dis := make([]int, n)
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[index] = 0
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		target := -1 // 寻找未访问的距离起点最近点
		for j := 0; j < n; j++ {
			if visited[j] == false && (target == -1 || dis[j] < dis[target]) {
				target = j
			}
		}
		visited[target] = true
		for j := 0; j < n; j++ { // 更新距离
			dis[j] = min(dis[j], dis[target]+arr[target][j])
		}
	}
	return dis
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

- 堆实现
- 注意，无向图和有向图

```go
// Dijkstra-堆优化：index => 其它地址距离
// arr => 邻接表 arr[a] = append(arr[a], [2]int{b, c}) 其中a=>b c
func Dijkstra(arr [][][2]int, index int) []int {
	n := len(arr)
	maxValue := math.MaxInt32
	dis := make([]int, n) // k到其他点的距离
	for i := 0; i < n; i++ {
		dis[i] = maxValue
	}
	dis[index] = 0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, [2]int{index, 0}) // 下标+距离
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([2]int) // 距离起点最近的点
		a := node[0]
		if dis[a] < node[1] { // 大于最短距离，跳过
			continue
		}
		for i := 0; i < len(arr[a]); i++ {
			b, c := arr[a][i][0], arr[a][i][1]
			if dis[a]+c < dis[b] { // 更新距离
				dis[b] = dis[a] + c
				heap.Push(&intHeap, [2]int{b, dis[b]})
			}
		}
	}
	return dis
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type IntHeap [][2]int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
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

| Title                                                                                       | Tag                                        | 难度     | 完成情况 |
| :---------------------------------------------------------------------------------------------| :--------------------------------------------| :--------| :------|
| [743.网络延迟时间](https://leetcode.cn/problems/network-delay-time/)                          | 堆、深度优先搜索、<br />广度优先搜索、图                    | Medium | 完成   |
| [1514.概率最大的路径](https://leetcode.cn/problems/path-with-maximum-probability/)             | 图                                          | Medium | 完成   |
| [1631.最小体力消耗路径](https://leetcode.cn/problems/path-with-minimum-effort/)                 | 深度优先搜索、广度优先搜索、并查集、<br />数组、二分查找、矩阵、堆（优先队列） | Medium | 完成   |
| [1976.到达目的地的方案数](https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/) | 图、拓扑排序、动态规划、最短路                            | Medium | 完成   |
| [2039.网络空闲的时刻](https://leetcode.cn/problems/the-time-when-the-network-becomes-idle/)    |                                            | Medium | 完成   |