package main

import "fmt"

func main() {
	first := ListNode{Val:1}
	firsttwo := ListNode{Val:2}
	firstthree := ListNode{Val:2}
	first.Next = &firsttwo
	firsttwo.Next = &firstthree
	deleteDuplicates(&first)

	for {
		fmt.Println(first.Val)
		if first.Next == nil{
			break
		}
		first = *first.Next
	}
}


type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil{
		return nil
	}
	temp := head
	for temp.Next != nil{
		if temp.Val == temp.Next.Val {
			temp.Next = temp.Next.Next
		}else {
			temp = temp.Next
		}
	}
	return head
}