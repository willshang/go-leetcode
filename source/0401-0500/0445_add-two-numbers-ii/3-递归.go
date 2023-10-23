package main

import "fmt"

func main() {
	first := ListNode{Val: 2}
	first1 := ListNode{Val: 4}
	first2 := ListNode{Val: 3}
	first.Next = &first1
	first1.Next = &first2

	second := ListNode{Val: 5}
	second1 := ListNode{Val: 6}
	second2 := ListNode{Val: 4}
	second.Next = &second1
	second1.Next = &second2

	fmt.Println(addTwoNumbers(&first, &second))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	a, b := l1, l2
	length1, length2 := 0, 0
	for a != nil {
		length1++
		a = a.Next
	}
	for b != nil {
		length2++
		b = b.Next
	}
	res, carry := add(l1, l2, length1, length2)
	if carry > 0 {
		return &ListNode{Val: carry, Next: res}
	}
	return res
}

func add(l1, l2 *ListNode, length1, length2 int) (res *ListNode, carry int) {
	if l1 != nil && l2 != nil {
		if l1.Next == nil && l2.Next == nil {
			val := l1.Val + l2.Val
			carry = val / 10
			res = &ListNode{Val: val % 10, Next: nil}
			return
		}
	}
	a := &ListNode{}
	var b, n int
	if length1 > length2 {
		a, b = add(l1.Next, l2, length1-1, length2)
		n = l1.Val + b
	} else if length1 < length2 {
		a, b = add(l1, l2.Next, length1, length2-1)
		n = l2.Val + b
	} else {
		a, b = add(l1.Next, l2.Next, length1-1, length2-1)
		n = l1.Val + l2.Val + b
	}
	res = &ListNode{Val: n % 10, Next: a}
	carry = n / 10
	return
}
