package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	num []int
}

type Node struct {
	root  *TreeNode
	index int
}

func Constructor(root *TreeNode) FindElements {
	if root == nil {
		return FindElements{num: []int{}}
	}
	num := make([]int, 0)
	num = append(num, 0)
	queue := make([]Node, 0)
	queue = append(queue, Node{
		root:  root,
		index: 0,
	})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node.root.Left != nil {
			temp := Node{
				root:  node.root.Left,
				index: node.index*2 + 1,
			}
			num = append(num, node.index*2+1)
			queue = append(queue, temp)
		}
		if node.root.Right != nil {
			temp := Node{
				root:  node.root.Right,
				index: node.index*2 + 2,
			}
			num = append(num, node.index*2+2)
			queue = append(queue, temp)
		}
	}
	return FindElements{num: num}
}

func (this *FindElements) Find(target int) bool {
	for i := 0; i < len(this.num); i++ {
		if this.num[i] == target {
			return true
		}
	}
	return false
}
