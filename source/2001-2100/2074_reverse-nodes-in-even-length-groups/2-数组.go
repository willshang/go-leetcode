package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetcode2074_反转偶数长度组的节点
func reverseEvenLengthGroups(head *ListNode) *ListNode {
	count := 1
	arr := make([]*ListNode, 0)
	for cur := head; cur != nil; cur = cur.Next {
		arr = append(arr, cur)
		if len(arr) == count || cur.Next == nil {
			if len(arr)%2 == 0 {
				for i := 0; i < len(arr)/2; i++ { // 交换值
					arr[i].Val, arr[len(arr)-1-i].Val = arr[len(arr)-1-i].Val, arr[i].Val
				}
			}
			arr = make([]*ListNode, 0)
			count++
		}
	}
	return head
}
