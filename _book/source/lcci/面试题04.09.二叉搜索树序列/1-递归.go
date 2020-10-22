package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 程序员面试金典04.09_二叉搜索树序列
var res [][]int

func BSTSequences(root *TreeNode) [][]int {
	res = make([][]int, 0)
	if root == nil {
		res = append(res, []int{})
		return res
	}
	dfs(append([]*TreeNode{}, root), make([]int, 0))
	return res
}

func dfs(arr []*TreeNode, path []int) {
	if len(arr) == 0 {
		res = append(res, path)
	}
	for i, node := range arr {
		temp := make([]int, len(path))
		copy(temp, path)
		temp = append(temp, node.Val)
		tempNode := make([]*TreeNode, len(arr))
		copy(tempNode, arr)
		tempNode = append(tempNode[:i], tempNode[i+1:]...) // 去除当前用过的
		if node.Left != nil {
			tempNode = append(tempNode, node.Left)
		}
		if node.Right != nil {
			tempNode = append(tempNode, node.Right)
		}
		dfs(tempNode, temp)
	}
}
