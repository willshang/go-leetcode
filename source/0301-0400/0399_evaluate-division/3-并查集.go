package main

func main() {

}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	m := make(map[string]int) // 计算对应的id
	for i := 0; i < len(equations); i++ {
		a, b := equations[i][0], equations[i][1]
		if _, ok := m[a]; ok == false {
			m[a] = len(m)
		}
		if _, ok := m[b]; ok == false {
			m[b] = len(m)
		}
	}
	fa, rank = Init(len(m))
	for i := 0; i < len(equations); i++ {
		a, b := m[equations[i][0]], m[equations[i][1]]
		union(a, b, values[i])
	}
	res := make([]float64, len(queries))
	for i := 0; i < len(queries); i++ {
		a, okA := m[queries[i][0]]
		b, okB := m[queries[i][1]]
		if okA == true && okB == true && find(a) == find(b) {
			res[i] = rank[a] / rank[b]
		} else {
			res[i] = -1
		}
	}
	return res
}

var fa []int
var rank []float64

// 初始化
func Init(n int) ([]int, []float64) {
	arr := make([]int, n)
	r := make([]float64, n)
	for i := 0; i < n; i++ {
		arr[i] = i
		r[i] = 1
	}
	return arr, r
}

// 查询
func find(x int) int {
	// 彻底路径压缩
	if fa[x] != x {
		origin := fa[x]
		fa[x] = find(fa[x])
		rank[x] = rank[x] * rank[origin] // 秩处理是难点
	}
	return fa[x]
}

// 合并
func union(i, j int, value float64) {
	x, y := find(i), find(j)
	rank[x] = value * rank[j] / rank[i] // 秩处理是难点
	fa[x] = y
}
