# Floyd

- 参考：多源最短路

## 0、定义

- 时间复杂度：O(n^3)
- 求：从任意起点出发，到达任意起点的最短距离
- 实现：邻接矩阵

## 1、操作

## 2、Go实现

```go
package main

func main() {

}

// Floyd 算法
func Floyd(arr [][]int) [][]int {
	n := len(arr)
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if arr[i][k]+arr[k][j] < arr[i][j] {
					arr[i][j] = arr[i][k] + arr[k][j]
				}
			}
		}
	}
	return arr
}
```

## 3、Leetcode

| Title                                                                                                                              | Tag                        | 难度     | 完成情况 |
| :------------------------------------------------------------------------------------------------------------------------------------| :----------------------------| :--------| :------|
| [399.除法求值](https://leetcode.cn/problems/evaluate-division/)                                                                    | 并查集、图                      | Medium | 完成   |
| [743.网络延迟时间](https://leetcode.cn/problems/network-delay-time/)                                                                 | 堆、深度优先搜索、<br />广度优先搜索、图    | Medium | 完成   |
| [1334.阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) | 图、动态规划、最短路                 | Medium | 完成   |
| [1462.课程表IV](https://leetcode.cn/problems/course-schedule-iv/)                                                                 | 深度优先搜索、广度优先搜索、<br />图、拓扑排序 | Medium | 完成   |
| [1976.到达目的地的方案数](https://leetcode.cn/problems/number-of-ways-to-arrive-at-destination/)                                        | 图、拓扑排序、动态规划、最短路            | Medium | 完成   |

