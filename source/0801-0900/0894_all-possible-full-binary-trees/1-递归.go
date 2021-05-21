package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var res map[int][]*TreeNode

func allPossibleFBT(n int) []*TreeNode {
	res = make(map[int][]*TreeNode)
	dfs(n)
	return res[n]
}

func dfs(n int) []*TreeNode {
	if _, ok := res[n]; !ok {
		arr := make([]*TreeNode, 0)
		if n == 1 {
			arr = append(arr, &TreeNode{Val: 0})
		} else if n%2 == 1 { // N只为奇数
			for i := 0; i < n; i++ {
				j := n - 1 - i
				for _, left := range dfs(i) {
					for _, right := range dfs(j) {
						root := &TreeNode{Val: 0}
						root.Left = left
						root.Right = right
						arr = append(arr, root)
					}
				}
			}
		}
		res[n] = arr
	}
	return res[n]
}
