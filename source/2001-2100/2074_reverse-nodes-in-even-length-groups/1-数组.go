package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	start := 1
	count := 2
	for start < len(arr) {
		end := start + count - 1 // 计算结尾下标
		if end >= len(arr) {
			end = len(arr) - 1
		}
		if (start-end+1)%2 == 0 { // 偶数个才反转
			reverse(arr, start, end) // 反转
		}
		start = start + count
		count = count + 1
	}
	temp := &ListNode{}
	node := temp
	for i := 0; i < len(arr); i++ {
		node.Next = &ListNode{Val: arr[i]}
		node = node.Next
	}
	return temp.Next
}

func reverse(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
