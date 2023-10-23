package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1110_删点成林
var res []*TreeNode
var m map[int]bool

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	res = make([]*TreeNode, 0)
	m = make(map[int]bool)
	for i := 0; i < len(to_delete); i++ {
		m[to_delete[i]] = true
	}
	root = dfs(root)
	if root != nil {
		res = append(res, root)
	}
	return res
}

func dfs(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = dfs(root.Left)
	root.Right = dfs(root.Right)
	if m[root.Val] == true {
		if root.Left != nil {
			res = append(res, root.Left)
		}
		if root.Right != nil {
			res = append(res, root.Right)
		}
		return nil
	}
	return root
}
