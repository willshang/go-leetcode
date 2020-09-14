package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

var res [][]int

func levelOrder(root *Node) [][]int {
	res = make([][]int, 0)
	if root == nil {
		return res
	}
	dfs(root, 0)
	return res
}

func dfs(root *Node, level int) {
	if root == nil {
		return
	}
	if len(res) == level {
		res = append(res, make([]int, 0))
	}
	res[level] = append(res[level], root.Val)
	for i := 0; i < len(root.Children); i++ {
		dfs(root.Children[i], level+1)
	}
}
