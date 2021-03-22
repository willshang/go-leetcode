package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode1609_奇偶树
func isEvenOddTree(root *TreeNode) bool {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	level := 1
	for len(queue) > 0 {
		length := len(queue)
		temp := make([]int, 0)
		for i := 0; i < length; i++ {
			temp = append(temp, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if level%2 == 0 {
			for j := 0; j < len(temp)/2; j++ {
				temp[j], temp[len(temp)-1-j] = temp[len(temp)-1-j], temp[j]
			}
		}
		for i := 0; i < len(temp); i++ {
			if i < len(temp)-1 && temp[i] >= temp[i+1] {
				return false
			}
			if temp[i]%2 != level%2 {
				return false
			}
		}
		queue = queue[length:]
		level++
	}
	return true
}
