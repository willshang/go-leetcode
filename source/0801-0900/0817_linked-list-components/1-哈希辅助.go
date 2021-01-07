package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode817_链表组件
func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(G); i++ {
		m[G[i]] = true
	}
	res := 0
	for head != nil {
		if m[head.Val] == true &&
			(head.Next == nil || m[head.Next.Val] == false) {
			res++
		}
		head = head.Next
	}
	return res
}
