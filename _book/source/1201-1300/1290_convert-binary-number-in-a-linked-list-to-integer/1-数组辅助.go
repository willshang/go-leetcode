package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	res := 0
	for i := 0; i < len(arr); i++ {
		res = res * 2
		res = res + arr[i]
	}
	return res
}
