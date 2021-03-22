package main

func main() {

}

type ListNode struct {
	Next  *ListNode
	Value int
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	res := &ListNode{}
	temp := res
	for i := 0; i < a; i++ {
		temp.Next = list1
		temp = temp.Next
		list1 = list1.Next
	}
	for list2 != nil {
		temp.Next = list2
		temp = temp.Next
		list2 = list2.Next
	}
	for i := a; i <= b; i++ {
		list1 = list1.Next
	}
	for list1 != nil {
		temp.Next = list1
		temp = temp.Next
		list1 = list1.Next
	}
	return res.Next
}
