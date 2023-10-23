package main

import "fmt"

func main() {
	a := Constructor()
	fmt.Println(a.Book(10, 20))
}

type MyCalendarThree struct {
	root *Node
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{root: &Node{
		start: 0,
		end:   1000000000,
	}}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	return this.root.Insert(start, end)
}

type Node struct {
	start int
	end   int
	count int
	delay int // 延迟更新线段树
	left  *Node
	right *Node
}

func (root *Node) getMid() int {
	return (root.start + root.end) / 2
}

func (root *Node) Left() *Node {
	if root.left == nil {
		root.left = &Node{
			start: root.start,
			end:   root.getMid(),
		}
	}
	return root.left
}

func (root *Node) Right() *Node {
	if root.right == nil {
		root.right = &Node{
			start: root.getMid(),
			end:   root.end,
		}
	}
	return root.right
}

func (root *Node) Insert(s, e int) int {
	if s <= root.start && root.end <= e { // 包含
		root.delay++
		root.count++
	} else if s < root.end && root.start < e { // 相交
		// 自上向下延迟更新
		root.Left().count = root.Left().count + root.delay
		root.Left().delay = root.Left().delay + root.delay
		root.Right().count = root.Right().count + root.delay
		root.Right().delay = root.Right().delay + root.delay
		root.delay = 0
		a := root.Left().Insert(s, e)
		b := root.Right().Insert(s, e)
		root.count = max(root.count, max(a, b))
	}
	return root.count
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
