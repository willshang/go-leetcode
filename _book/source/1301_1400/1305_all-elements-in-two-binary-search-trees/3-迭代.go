package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	res := make([]int, 0)
	arr1 := inorderTraversal(root1)
	arr2 := inorderTraversal(root2)
	i, j := 0, 0
	for i < len(arr1) || j < len(arr2) {
		if i < len(arr1) && (j == len(arr2) || arr1[i] < arr2[j]) {
			res = append(res, arr1[i])
			i++
		} else {
			res = append(res, arr2[j])
			j++
		}
	}
	return res
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		last := len(stack) - 1
		res = append(res, stack[last].Val)
		root = stack[last].Right
		stack = stack[:last]
	}
	return res
}
