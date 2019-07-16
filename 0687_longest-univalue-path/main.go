package main

import "fmt"

func main() {
	root := TreeNode{Val: 5}
	rootLeft := TreeNode{Val: 4}
	rootRight := TreeNode{Val: 5}
	rootLeftLeft := TreeNode{Val: 1}
	rootLeftRight := TreeNode{Val: 1}
	rootRightRight := TreeNode{Val: 5}

	root.Left = &rootLeft
	root.Right = &rootRight
	rootLeft.Left = &rootLeftLeft
	rootLeft.Right = &rootLeftRight
	rootRight.Right = &rootRightRight

	fmt.Println(longestUnivaluePath(&root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root, &maxLen)
	return maxLen
}

func helper(n *TreeNode, maxLen *int) int {
	if n == nil {
		return 0
	}
	l := helper(n.Left, maxLen)
	r := helper(n.Right, maxLen)

	if n.Left != nil && n.Val == n.Left.Val {
		l++
	} else {
		l = 0
	}

	if n.Right != nil && n.Val == n.Right.Val {
		r++
	} else {
		r = 0
	}
	if l+r > *maxLen {
		*maxLen = l + r
	}
	return max(l, r)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0
	helper(root,&maxLen)
	return maxLen
}

func helper(n *TreeNode, maxLen *int)int  {
	if n == nil{
		return 0
	}
	l := helper(n.Left,maxLen)
	r := helper(n.Right,maxLen)
	res := 0

	if n.Left != nil && n.Val == n.Left.Val{
		*maxLen = max(*maxLen,l+1)
		res = max(res,l+1)
	}

	if n.Right != nil && n.Val == n.Right.Val{
		*maxLen = max(*maxLen,r+1)
		res = max(res,r+1)
	}
	if n.Left != nil && n.Val == n.Left.Val && n.Right != nil && n.Val == n.Right.Val{
		*maxLen = max(*maxLen, l+r+2)
	}
	return res
}

func max(a, b int)int  {
	if a > b {
		return a
	}
	return b
}
*/
