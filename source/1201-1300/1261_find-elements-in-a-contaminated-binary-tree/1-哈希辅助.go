package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1261_在受污染的二叉树中查找元素
type FindElements struct {
	m map[int]bool
}

type Node struct {
	root  *TreeNode
	index int
}

func Constructor(root *TreeNode) FindElements {
	if root == nil {
		return FindElements{m: map[int]bool{}}
	}
	m := make(map[int]bool)
	m[0] = true
	queue := make([]Node, 0)
	queue = append(queue, Node{root: root, index: 0})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		var temp Node
		if node.root.Left != nil {
			temp = Node{
				root:  node.root.Left,
				index: node.index*2 + 1,
			}
			m[node.index*2+1] = true
		}
		if node.root.Right != nil {
			temp = Node{
				root:  node.root.Right,
				index: node.index*2 + 2,
			}
			m[node.index*2+2] = true
		}
		queue = append(queue, temp)
	}
	return FindElements{m: m}
}

func (this *FindElements) Find(target int) bool {
	return this.m[target]
}
