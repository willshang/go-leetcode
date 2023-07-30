# 图-Kruskal

- 参考：

## 0、定义

- 适合：稀疏图

## 1、操作

- 排序
- 并查集判断

## 2、Go实现

- 稀疏图

```go
package main

func main() {

}

// Kruskal 算法
func Kruskal(n int, arr [][3]int) int {
	res := 0
	fa = Init(n)
	count := 0
	for i := 0; i < len(arr); i++ {
		a, b, c := arr[i][0], arr[i][1], arr[i][2]
		if query(a, b) == false {
			union(a, b)
			res = res + c
			count++
			if count == n-1 {
				break
			}
		}
	}
	return res
}

var fa []int

// Init 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	return arr
}

// 查询
func find(x int) int {
	if fa[x] != x {
		fa[x] = find(fa[x])
	}
	return fa[x]
}

// 合并
func union(i, j int) {
	fa[find(i)] = find(j)
}

func query(i, j int) bool {
	return find(i) == find(j)
}
```

## 3、Leetcode

| Title                                                                           | Tag          | 难度     | 完成情况 |
|:--------------------------------------------------------------------------------|:-------------|:-------|:-----|
| [1584.连接所有点的最小费用](https://leetcode.cn/problems/min-cost-to-connect-all-points/) | 并查集、数组、最小生成树 | Medium | 完成   |