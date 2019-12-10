package main

import "fmt"

func main() {
	root := TreeNode{Val: 1}
	rootfirst := TreeNode{Val: 2}
	rootSecond := TreeNode{Val: 3}
	root.Left = &rootfirst
	root.Right = &rootSecond

	second := TreeNode{Val: 1}
	secondfirst := TreeNode{Val: 2}
	secondSecond := TreeNode{Val: 4}
	second.Left = &secondfirst
	second.Right = &secondSecond

	fmt.Println(isSameTree(&root, &second))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	var queueP, queueQ []*TreeNode
	if p != nil {
		queueP = append(queueP, p)
		queueQ = append(queueQ, q)
	}

	for len(queueP) > 0 && len(queueQ) > 0 {
		tempP := queueP[0]
		queueP = queueP[1:]

		tempQ := queueQ[0]
		queueQ = queueQ[1:]

		if tempP.Val != tempQ.Val {
			return false
		}

		if (tempP.Left != nil && tempQ.Left == nil) ||
			(tempP.Left == nil && tempQ.Left != nil) {
			return false
		}
		if tempP.Left != nil {
			queueP = append(queueP, tempP.Left)
			queueQ = append(queueQ, tempQ.Left)
		}

		if (tempP.Right != nil && tempQ.Right == nil) ||
			(tempP.Right == nil && tempQ.Right != nil) {
			return false
		}
		if tempP.Right != nil {
			queueP = append(queueP, tempP.Right)
			queueQ = append(queueQ, tempQ.Right)
		}
	}
	return true
}
