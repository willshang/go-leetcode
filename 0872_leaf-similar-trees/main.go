package main

import "fmt"

func main() {
	first := TreeNode{Val: 1}
	second := TreeNode{Val: 2}
	third := TreeNode{Val: 3}
	first.Left = &second
	first.Right = &third

	first1 := TreeNode{Val: 1}
	second1 := TreeNode{Val: 2}
	third1 := TreeNode{Val: 3}
	first1.Left = &second1
	first1.Right = &third1

	fmt.Println(leafSimilar(&first, &first1))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	a1 := [100]int{}
	s1 := a1[:0]
	search(root1, &s1)

	a2 := [100]int{}
	s2 := a2[:0]
	search(root2, &s2)

	return a1 == a2
}

func search(root *TreeNode, sp *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*sp = append(*sp, root.Val)
		return
	}
	search(root.Left, sp)
	search(root.Right, sp)
}
