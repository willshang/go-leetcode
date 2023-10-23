package main

func main() {

}

func minSwapsCouples(row []int) int {
	n := len(row) / 2
	fa := make([]int, n)
	for i := 0; i < n; i++ {
		fa[i] = i
	}
	// 将每张沙发上的两个人员编号union一下，如果本来编号就相同，则表示两个人是一类
	for i := 0; i < len(row); i = i + 2 {
		a, b := row[i]/2, row[i+1]/2
		union(fa, a, b)
	}
	res := 0
	for i := 0; i < n; i++ {
		// 几个相同，就有几个环
		if find(fa, i) == i {
			res++
		}
	}
	// 如果3组1个环，需要的次数是3-1=2，另外4组1个环，需要的次数是4-1=3=》4+3-2=5
	// 组数-减去环数
	return n - res
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
