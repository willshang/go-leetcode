package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	flag := false
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
		if flag == true {
			for j := 0; j < len(temp)/2; j++ {
				temp[j], temp[len(temp)-1-j] = temp[len(temp)-1-j], temp[j]
			}
			flag = false
		} else {
			flag = true
		}
		if len(temp) != level {
			return false
		}
		if temp[0]%2 != level%2 {
			return false
		}
		for j := 1; j < len(temp); j++ {
			if temp[j]%2 != level%2 || temp[j-1] >= temp[j] {
				return false
			}
		}
		queue = queue[length:]
		level++
	}
	return true
}
