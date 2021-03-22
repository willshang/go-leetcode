package main

func main() {

}

type ListNode struct {
	Next  *ListNode
	Value int
}

// leetcode1669_合并两个链表
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cur := list1
	for i := 0; i < a-1; i++ {
		cur = cur.Next
	}
	temp := cur.Next
	for i := 0; i < (b - a + 1); i++ {
		temp = temp.Next
	}
	cur.Next = list2
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = temp
	return list1
}
