package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseEvenLengthGroups(head *ListNode) *ListNode {
	count := 1
	cur := head
	prev := &ListNode{}
	for cur != nil {
		c := 0
		temp := cur
		for c < count && temp != nil {
			c++
			temp = temp.Next
		}
		if c%2 == 1 {
			for i := 0; i < c; i++ {
				prev, cur = cur, cur.Next // 指针后移
			}
		} else { // 反转链表
			for i := 0; i < c-1; i++ {
				prev.Next, cur.Next.Next, cur.Next = cur.Next, prev.Next, cur.Next.Next
			}
			prev, cur = cur, cur.Next
		}
		count++
	}
	return head
}
