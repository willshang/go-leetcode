package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode894_所有可能的满二叉树
func allPossibleFBT(n int) []*TreeNode {
	res := make(map[int][]*TreeNode)
	res[1] = []*TreeNode{&TreeNode{Val: 0}}
	for index := 3; index <= n; index = index + 2 {
		for i := 1; i < index; i = i + 2 {
			j := index - 1 - i
			for _, left := range res[i] {
				for _, right := range res[j] {
					res[index] = append(res[index], &TreeNode{
						Val:   0,
						Left:  left,
						Right: right,
					})
				}
			}
		}
	}
	return res[n]
}
