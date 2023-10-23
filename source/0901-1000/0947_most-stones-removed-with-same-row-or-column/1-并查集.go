package main

func main() {

}

// leetcode947_移除最多的同行或同列石头
func removeStones(stones [][]int) int {
	fa := make([]int, 20000)
	for i := 0; i < 20000; i++ {
		fa[i] = i
	}
	for i := 0; i < len(stones); i++ {
		a, b := stones[i][0], stones[i][1]
		// 连接同一行、列
		union(fa, a, b+10000)
	}
	m := make(map[int]bool)
	for i := 0; i < len(stones); i++ {
		a := stones[i][0]
		m[find(fa, a)] = true
	}
	return len(stones) - len(m)
}

func union(fa []int, a, b int) {
	fa[find(fa, a)] = find(fa, b)
}

func find(fa []int, a int) int {
	for fa[a] != a {
		fa[a] = fa[fa[a]]
		a = fa[a]
	}
	return a
}
