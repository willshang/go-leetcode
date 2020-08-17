package main

func main() {
	first := ListNode{Val: 3}
	second := ListNode{Val: 2}
	third := ListNode{Val: 0}
	fourth := ListNode{Val: -4}
	first.Next = &second
	second.Next = &third
	third.Next = &fourth
	fourth.Next = &second
	reorderList(&first)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	cur := head
	arr := make([]*ListNode, 0)
	for cur != nil {
		arr = append(arr, cur)
		cur = cur.Next
	}
	res := make([]*ListNode, 0)
	for i := 0; i < len(arr)/2; i++ {
		res = append(res, arr[i], arr[len(arr)-1-i])
	}
	if len(arr)%2 == 1 {
		res = append(res, arr[len(arr)/2])
	}
	cur = head
	for i := 1; i < len(res); i++ {
		cur.Next = res[i]
		cur = cur.Next
	}
	cur.Next = nil
}
