package main

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

// leetcode429_N叉树的层序遍历
func levelOrder(root *Node) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		temp := make([]int, 0)
		length := len(queue)
		for i := 0; i < length; i++ {
			if queue[i] != nil {
				temp = append(temp, queue[i].Val)
				for j := 0; j < len(queue[i].Children); j++ {
					queue = append(queue, queue[i].Children[j])
				}
			}
		}
		res = append(res, temp)
		queue = queue[length:]
	}
	return res
}
