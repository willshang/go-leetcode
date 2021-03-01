package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	res := &ListNode{Next: head}
	for cur := res; cur != nil; {
		sum := 0
		for temp := cur.Next; temp != nil; temp = temp.Next {
			sum = sum + temp.Val
			if sum == 0 {
				cur.Next = temp.Next // 调整指针
			}
		}
		cur = cur.Next
	}
	return res.Next
}
