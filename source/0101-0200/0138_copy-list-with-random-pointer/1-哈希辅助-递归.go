package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var m map[*Node]*Node

func copyRandomList(head *Node) *Node {
	m = make(map[*Node]*Node)
	return copyList(head)
}

func copyList(head *Node) *Node {
	if head == nil {
		return head
	}
	if node, ok := m[head]; ok {
		return node
	}
	temp := &Node{
		Val:    head.Val,
		Next:   nil,
		Random: nil,
	}
	m[head] = temp
	temp.Next = copyList(head.Next)
	temp.Random = copyList(head.Random)
	return temp
}
