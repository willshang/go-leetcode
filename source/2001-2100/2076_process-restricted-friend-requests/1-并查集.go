package main

func main() {

}

func friendRequests(n int, restrictions [][]int, requests [][]int) []bool {
	fa = Init(n)
	res := make([]bool, len(requests))
	for i := 0; i < len(requests); i++ {
		a, b := requests[i][0], requests[i][1]
		x, y := find(a), find(b)
		if x == y {
			res[i] = true
		} else {
			flag := true
			for j := 0; j < len(restrictions); j++ { // 尝试每个限制条件
				c, d := restrictions[j][0], restrictions[j][1]
				u, v := find(c), find(d)
				if (x == u && y == v) || (x == v && y == u) { // 有限制
					flag = false
					break
				}
			}
			if flag == true {
				res[i] = true
				union(x, y)
			}
		}
	}
	return res
}

var fa []int

// 初始化
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
