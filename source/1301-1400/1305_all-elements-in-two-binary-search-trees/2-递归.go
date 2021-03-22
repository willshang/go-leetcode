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
	arr1, arr2 := make([]int, 0), make([]int, 0)
	dfs(root1, &arr1)
	dfs(root2, &arr2)
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

func dfs(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, arr)
	*arr = append(*arr, root.Val)
	dfs(root.Right, arr)
}
