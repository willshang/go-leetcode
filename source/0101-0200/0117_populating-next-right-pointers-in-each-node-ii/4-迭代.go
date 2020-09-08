package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// leetcode117_填充节点下一个右侧节点指针II
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		down := &Node{}
		prev := down
		for cur != nil {
			if cur.Left != nil {
				prev.Next = cur.Left
				prev = prev.Next
			}
			if cur.Right != nil {
				prev.Next = cur.Right
				prev = prev.Next
			}
			cur = cur.Next // 当前层级移动
		}
		cur = down.Next // 移到下一层最左边
	}
	return root
}
