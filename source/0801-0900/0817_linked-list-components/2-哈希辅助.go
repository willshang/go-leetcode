package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func numComponents(head *ListNode, G []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(G); i++ {
		m[G[i]] = true
	}
	res := 0
	flag := false
	for head != nil {
		if m[head.Val] == true {
			if flag == false {
				res++
				flag = true
			}
		} else {
			flag = false
		}
		head = head.Next
	}
	return res
}
