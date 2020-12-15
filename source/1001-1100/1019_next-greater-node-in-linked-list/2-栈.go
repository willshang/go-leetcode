package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode1019_链表中的下一个更大节点
func nextLargerNodes(head *ListNode) []int {
	arr := make([]int, 0)
	if head == nil {
		return arr
	}
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := make([]int, len(arr))
	stack := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		for len(stack) > 0 && arr[i] > arr[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			res[last] = arr[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
