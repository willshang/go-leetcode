package main

import "fmt"

func main() {
	n := 200
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	union(fa, 10, 20)
	union(fa, 10, 30)
	fmt.Println(find(fa, 20), find(fa, 30))
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
