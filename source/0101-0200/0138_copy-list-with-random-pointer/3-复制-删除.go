package main

func main() {

}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	res := copyNext(head)
	res = copyRandom(res)
	res = cutEven(res)
	return res
}

// 原1-复制1-原2-复制2
func copyNext(head *Node) *Node {
	p := head
	for p != nil {
		node := new(Node)
		node.Val = p.Val
		node.Next = p.Next
		p.Next = node
		p = node.Next
	}
	return head
}

func copyRandom(head *Node) *Node {
	p := head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	return head
}

func cutEven(head *Node) *Node {
	oldNode := head
	newNode := head.Next
	cur := newNode
	for oldNode != nil {
		oldNode.Next = oldNode.Next.Next
		if newNode.Next != nil {
			newNode.Next = newNode.Next.Next
		}
		oldNode = oldNode.Next
		newNode = newNode.Next
	}
	return cur
}
