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
	for cur.Left != nil {
		parent := cur
		for parent != nil {
			parent.Left.Next = parent.Right // 左节点连接右节点
			if parent.Next != nil {
				// 图中的5->6 ，左子树的右节点->右子树的左节点
				parent.Right.Next = parent.Next.Left
			}
			parent = parent.Next
		}
		cur = cur.Left // 移到下一层最左边
	}
	return root
}
