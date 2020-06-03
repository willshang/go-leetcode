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
	temp := new(Node)
	temp.Next = head
	p := head
	for p != nil {
		node := new(Node)
		node.Val = p.Val
		node.Next = p.Next
		p.Next = node
		p = p.Next.Next
	}
	return temp.Next
}

func copyRandom(head *Node) *Node {
	temp := new(Node)
	temp.Next = head
	p := head
	for p != nil {
		if p.Random != nil {
			p.Next.Random = p.Random.Next
		}
		p = p.Next.Next
	}
	return temp.Next
}

func cutEven(head *Node) *Node {
	temp := new(Node)
	res := head
	p := head
	for p != nil {
		res.Next = p.Next
		p.Next = p.Next.Next
		p = p.Next
		res = res.Next
	}
	return temp.Next
}
