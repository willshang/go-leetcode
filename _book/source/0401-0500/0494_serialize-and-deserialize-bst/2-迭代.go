package main

import (
	"strconv"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node != nil {
			res = append(res, strconv.Itoa(node.Val))
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, "#")
		}
	}
	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 || data == "" {
		return nil
	}
	res := strings.Split(data, ",")
	root := &TreeNode{}
	root.Val, _ = strconv.Atoi(res[0])
	res = res[1:]
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		if res[0] != "#" {
			left, _ := strconv.Atoi(res[0])
			queue[0].Left = &TreeNode{Val: left}
			queue = append(queue, queue[0].Left)
		}
		if res[1] != "#" {
			right, _ := strconv.Atoi(res[1])
			queue[0].Right = &TreeNode{Val: right}
			queue = append(queue, queue[0].Right)
		}
		queue = queue[1:]
		res = res[2:]
	}
	return root
}
