package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	stack := make([]*TreeNode, 0)
	var cur *TreeNode
	for i := 0; i < len(nums); i++ {
		cur = &TreeNode{
			Val: nums[i],
		}
		// 递减栈
		for len(stack) > 0 && stack[len(stack)-1].Val < cur.Val {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// top选择cur或者栈顶数据作为父节点
			if len(stack) > 0 && stack[len(stack)-1].Val < cur.Val {
				stack[len(stack)-1].Right = top
			} else {
				cur.Left = top
			}
		}
		stack = append(stack, cur)
	}
	// 没有右边节点，栈顶元素作为第二个栈顶元素的右节点
	for len(stack) > 0 {
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if len(stack) > 0 {
			stack[len(stack)-1].Right = cur
		}
	}
	return cur
}
