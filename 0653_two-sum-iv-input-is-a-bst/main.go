package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	if root == nil {
		return false
	}
	m := map[int]int{}
	return helper(root, k, m)
}

func helper(node *TreeNode, k int, m map[int]int) bool {
	if node == nil {
		return false
	}
	if _, ok := m[k-node.Val]; ok {
		return true
	}
	m[node.Val] = node.Val
	return helper(node.Left, k, m) || helper(node.Right, k, m)
}

/*func findTarget(root *TreeNode, k int) bool {
	return helper(root,root,k)
}

func helper(root, searchRoot *TreeNode, k int) bool  {
	if root == nil{
		return false
	}
	return (root.Val*2 != k && findNode(searchRoot, k-root.Val)) ||
		helper(root.Left,searchRoot,k) ||
		helper(root.Right,searchRoot,k)
}

func findNode(root *TreeNode, target int)bool  {
	if root == nil{
		return false
	}

	if root.Val == target{
		return true
	}

	if root.Val < target{
		return findNode(root.Right,target)
	}
	return findNode(root.Left,target)
}*/
