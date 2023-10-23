package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

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
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] < arr[j] {
				res[i] = arr[j]
				break
			}
		}
	}
	return res
}
