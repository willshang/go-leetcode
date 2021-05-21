package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	if n == 1 {
		return []*TreeNode{&TreeNode{Val: 0}}
	}
	res := make([]*TreeNode, 0)
	for i := 1; i < n; i = i + 2 {
		leftResult := allPossibleFBT(i)
		rightResult := allPossibleFBT(n - 1 - i)
		for _, left := range leftResult {
			for _, right := range rightResult {
				res = append(res, &TreeNode{
					Val:   0,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	return res
}
