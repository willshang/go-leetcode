package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	m := make(map[int]*ListNode)
	res := head
	sum := 0
	for cur := head; cur != nil; cur = cur.Next {
		sum = sum + cur.Val
		if sum == 0 { // 当前都为0
			res = cur.Next
			m = make(map[int]*ListNode)
			continue
		}
		if _, ok := m[sum]; ok == false {
			m[sum] = cur
			continue
		}
		// 出现重复
		first := m[sum]
		tempSum := sum
		for temp := first.Next; temp != cur; temp = temp.Next {
			tempSum = tempSum + temp.Val
			delete(m, tempSum)
		}
		first.Next = cur.Next
	}
	return res
}
