package main

func main() {

}

// 剑指OfferII117.相似的字符串
func numSimilarGroups(strs []string) int {
	n := len(strs)
	fa = Init(n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if judge(strs[i], strs[j]) == true { // 满足条件，连通
				union(i, j)
			}
		}
	}
	return count
}

func judge(a, b string) bool {
	if a == b {
		return true
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
			if count > 2 {
				return false
			}
		}
	}
	return true
}

var fa []int
var count int

// 初始化
func Init(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	count = n
	return arr
}

// 查询
func find(x int) int {
	if fa[x] == x {
		return x
	}
	// 路径压缩
	fa[x] = find(fa[x])
	return fa[x]
}

// 合并
func union(i, j int) {
	x, y := find(i), find(j)
	if x != y {
		fa[x] = y
		count--
	}
}
