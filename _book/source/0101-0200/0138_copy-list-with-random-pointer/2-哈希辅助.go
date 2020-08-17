package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// leetcode138_复制带随机指针的链表
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	res := new(Node)
	m := make(map[*Node]*Node)
	temp := head
	p := res
	for temp != nil {
		node := &Node{
			Val:    temp.Val,
			Next:   nil,
			Random: nil,
		}
		m[temp] = node
		p.Next = node
		p = p.Next
		temp = temp.Next

	}
	temp = head
	p = res.Next
	for temp != nil {
		p.Random = m[temp.Random]
		p = p.Next
		temp = temp.Next
	}
	return res.Next
}
