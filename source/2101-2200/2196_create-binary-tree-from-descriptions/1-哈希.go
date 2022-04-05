package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode2196_根据描述创建二叉树
func createBinaryTree(descriptions [][]int) *TreeNode {
	m := make(map[int]*TreeNode)
	visited := make(map[int]bool) // 保存子节点
	for i := 0; i < len(descriptions); i++ {
		a, b, c := descriptions[i][0], descriptions[i][1], descriptions[i][2]
		if m[a] == nil {
			m[a] = &TreeNode{Val: a}
		}
		if m[b] == nil {
			m[b] = &TreeNode{Val: b}
		}
		if c == 1 {
			m[a].Left = m[b]
		} else {
			m[a].Right = m[b]
		}
		visited[b] = true
	}
	for k := range m {
		if visited[k] == false { // 根节点不属于子节点
			return m[k]
		}
	}
	return nil
}
