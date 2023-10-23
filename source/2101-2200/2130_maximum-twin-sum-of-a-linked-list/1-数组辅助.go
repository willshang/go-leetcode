package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode2130_链表最大孪生和
func pairSum(head *ListNode) int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := 0
	for i := 0; i < len(arr)/2; i++ {
		res = max(res, arr[i]+arr[len(arr)-1-i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
