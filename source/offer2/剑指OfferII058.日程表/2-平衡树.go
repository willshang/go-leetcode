package main

func main() {

}

type MyCalendar struct {
	root *Node
}

func Constructor() MyCalendar {
	return MyCalendar{root: nil}
}

func (this *MyCalendar) Book(start int, end int) bool {
	node := &Node{
		start: start,
		end:   end,
		left:  nil,
		right: nil,
	}
	if this.root == nil {
		this.root = node
		return true
	}
	return this.root.Insert(node)
}

type Node struct {
	start int
	end   int
	left  *Node
	right *Node
}

func (this *Node) Insert(node *Node) bool {
	if node.start >= this.end {
		if this.right == nil {
			this.right = node
			return true
		}
		return this.right.Insert(node)
	} else if node.end <= this.start {
		if this.left == nil {
			this.left = node
			return true
		}
		return this.left.Insert(node)
	}
	return false
}
