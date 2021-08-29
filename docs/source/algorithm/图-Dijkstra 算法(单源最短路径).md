# 图-最短路径算法-Dijkstra

## 0、定义

- **解决问题**：从图中的某个顶点出发到达另外一个顶点的所经过的边的权重和最小的一条路径，称为最短路径。
- 算法：迪杰斯特拉算法主要特点是**从起始点开始**，采用**贪心算法的策略**，每次遍历到始点距离最近且未访问过的顶点的邻接节点，**直到扩展到终点为止。**

## 1、操作

## 2、Go实现

- 不加堆实现

```go

```

- 堆实现

```go

```

## 3、Leetcode

| Title                                                        | Tag                                                          | 难度   | 完成情况 |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------ | -------- |
| [743.网络延迟时间](https://leetcode-cn.com/problems/network-delay-time/) | 堆、深度优先搜索、<br />广度优先搜索、图                     | Medium | 完成     |
| [1514.概率最大的路径](https://leetcode-cn.com/problems/path-with-maximum-probability/) | 图                                                           | Medium | 完成     |
| [1631.最小体力消耗路径](https://leetcode-cn.com/problems/path-with-minimum-effort/) | 深度优先搜索、广度优先搜索、并查集、<br />数组、二分查找、矩阵、堆（优先队列） | Medium | 完成     |
| [1976.到达目的地的方案数](https://leetcode-cn.com/problems/number-of-ways-to-arrive-at-destination/) | 图、拓扑排序、动态规划、最短路                               | Medium | 完成     |