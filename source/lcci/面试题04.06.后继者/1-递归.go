package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 程序员面试金典04.06_后继者
var res []*TreeNode

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	res = make([]*TreeNode, 0)
	dfs(root)
	for i := 0; i < len(res)-1; i++ {
		if res[i] == p {
			return res[i+1]
		}
	}
	return nil
}

func dfs(root *TreeNode) {
	if root == nil {
		return
	}
	dfs(root.Left)
	res = append(res, root)
	dfs(root.Right)
}
