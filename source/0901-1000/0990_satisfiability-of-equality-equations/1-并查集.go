package main

func main() {

}

// leetcode990_等式方程的可满足性
func equationsPossible(equations []string) bool {
	fa := make([]int, 26)
	for i := 0; i < 26; i++ {
		fa[i] = i
	}
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '=' {
			a, b := int(equations[i][0]-'a'), int(equations[i][3]-'a')
			union(fa, a, b)
		}
	}
	for i := 0; i < len(equations); i++ {
		if equations[i][1] == '!' {
			a, b := int(equations[i][0]-'a'), int(equations[i][3]-'a')
			if find(fa, a) == find(fa, b) {
				return false
			}
		}
	}
	return true
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
