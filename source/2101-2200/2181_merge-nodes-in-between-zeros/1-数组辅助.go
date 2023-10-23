package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	temp := make([]int, 0)
	for head != nil {
		temp = append(temp, head.Val)
		head = head.Next
	}
	sum := 0
	arr := make([]int, 0)
	for i := 1; i < len(temp); i++ {
		if temp[i] == 0 && sum != 0 {
			arr = append(arr, sum)
			sum = 0
		} else {
			sum = sum + temp[i]
		}
	}
	res := &ListNode{}
	t := res
	for i := 0; i < len(arr); i++ {
		t.Next = &ListNode{
			Val:  arr[i],
			Next: nil,
		}
		t = t.Next
	}
	return res.Next
}
