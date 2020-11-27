package main

import "fmt"

func main() {

}

type Node struct {
	LeftValue  int
	RightValue int
	Left       *Node
	Right      *Node
	Value      int
}

// 1、build函数
func (this *Node) build(left, right int) *Node {
	this.LeftValue = left
	this.RightValue = right
	this.Value = 0
	mid := left + (right-left)/2
	if left+1 < right {
		if this.LeftValue != mid {
			newLeft := Node{}
			this.Left = newLeft.build(left, mid)
		}
		if this.RightValue != mid {
			newRight := &Node{}
			this.Right = newRight.build(mid, right)
		}
		return this
	}
	return this
}

// 2、insert函数
func (this *Node) Insert(left, right int) {
	mid := this.LeftValue + (this.RightValue-this.LeftValue)/2
	if left == this.LeftValue && right == this.RightValue {
		this.Value = 1
	} else if right <= mid {
		if this.Left != nil {
			this.Left.Insert(left, right)
		}
	} else if left >= mid {
		if this.Right != nil {
			this.Right.Insert(left, right)
		}
	} else {
		if this.Left != nil {
			this.Left.Insert(left, mid)
		}
		if this.Right != nil {
			this.Right.Insert(mid, right)
		}
	}
}

// 3.求和
func (this *Node) Sum() int {
	if this.Value == 1 {
		return this.RightValue - this.LeftValue
	} else if this.RightValue-this.LeftValue == 1 {
		return 0
	}
	return this.Left.Sum() - this.Right.Sum()
}

// 4.中序遍历
func (this *Node) DFS() {
	if this != nil {
		this.Left.DFS()
		fmt.Println(this.Value)
		this.Right.DFS()
	}
}
