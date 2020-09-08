package main

func main() {

}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}
	cur := root
	for cur != nil {
		var prev, down *Node
		for cur != nil {
			if cur.Left != nil {
				if prev != nil {
					prev.Next = cur.Left
				} else {
					down = cur.Left
				}
				prev = cur.Left
			}
			if cur.Right != nil {
				if prev != nil {
					prev.Next = cur.Right
				} else {
					down = cur.Right
				}
				prev = cur.Right
			}
			cur = cur.Next // 当前层级移动
		}
		cur = down // 移到下一层最左边
	}
	return root
}
