package main

func main() {

}

func countPairs(n int, edges [][]int) int64 {
	res := int64(0)
	fa = Init(n)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		union(a, b)
	}
	m := make(map[int]int)
	for i := 0; i < n; i++ {
		m[find(i)]++ // 统计每组个数
	}
	for _, v := range m {
		res = res + int64(v)*int64(n-v)
	}
	return res / 2
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
