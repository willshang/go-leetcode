package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// leetcode109_有序链表转换二叉搜索树
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	arr := make([]int, 0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}
	return sortArr(arr)
}

func sortArr(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   arr[len(arr)/2],
		Left:  sortArr(arr[:len(arr)/2]),
		Right: sortArr(arr[len(arr)/2+1:]),
	}
}
