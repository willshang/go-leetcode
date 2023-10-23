package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: v}
	}
	if d == 1 {
		return &TreeNode{Val: v, Left: root}
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	var level = 1
	for len(queue) > 0 {
		level++
		length := len(queue)
		if level == d {
			for i := 0; i < length; i++ {
				queue[i].Left = &TreeNode{
					Left: queue[i].Left,
					Val:  v,
				}
				queue[i].Right = &TreeNode{
					Right: queue[i].Right,
					Val:   v,
				}
			}
		}
		for i := 0; i < length; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[length:]
	}
	return root
}
