package main

func main()  {
	
}



type ListNode struct {
     Val int
     Next *ListNode
}


func hasCycle(head *ListNode) bool {
	m := make(map[*ListNode]bool)
	for head != nil{
		if m[head] {
			return true
		}
		m[head] = true
		head = head.Next
	}
	return false
}