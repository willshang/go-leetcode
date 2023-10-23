package main

func main() {

}

// leetcode1579_保证图可完全遍历
func maxNumEdgesToRemove(n int, edges [][]int) int {
	res := len(edges)
	alice, bob := &UnionFind{}, &UnionFind{}
	alice.Init(n)
	bob.Init(n)
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1]-1, edges[i][2]-1
		if a == 3 && (alice.query(b, c) == false || bob.query(b, c) == false) { // 公共边不在一个集合
			alice.union(b, c)
			bob.union(b, c)
			res--
		}
	}
	for i := 0; i < len(edges); i++ {
		a, b, c := edges[i][0], edges[i][1]-1, edges[i][2]-1
		if a == 1 && alice.query(b, c) == false { // 单边不在一个集合
			alice.union(b, c)
			res--
		}
		if a == 2 && bob.query(b, c) == false { // 单边不在一个集合
			bob.union(b, c)
			res--
		}
	}
	if alice.count > 1 || bob.count > 1 {
		return -1
	}
	return res
}

type UnionFind struct {
	fa    []int
	count int
}

// 初始化
func (u *UnionFind) Init(n int) {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i
	}
	u.count = n
	u.fa = arr
}

// 查询
func (u UnionFind) find(x int) int {
	if u.fa[x] == x {
		return x
	}
	// 路径压缩
	u.fa[x] = u.find(u.fa[x])
	return u.fa[x]
}

// 合并
func (u *UnionFind) union(i, j int) {
	x, y := u.find(i), u.find(j)
	if x != y {
		u.fa[x] = y
		u.count--
	}
}

func (u UnionFind) query(i, j int) bool {
	return u.find(i) == u.find(j)
}
