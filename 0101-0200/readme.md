# 0101-0200

[TOC]
## 101. 对称二叉树
### 题目
``` 
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树 [1,2,2,3,4,4,3] 是对称的。
    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:
    1
   / \
  2   2
   \   \
   3    3

说明:
如果你可以运用递归和迭代两种方法解决这个问题，会很加分。
```
### 解答思路
| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(n)       |
| 02   | 迭代 | O(n)       | O(n)       |

```go
// 递归
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recur(root.Left, root.Right)
}

func recur(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}

	return left.Val == right.Val &&
		recur(left.Left, right.Right) &&
		recur(left.Right, right.Left)
}
// 迭代
func isSymmetric(root *TreeNode) bool {
	leftQ := make([]*TreeNode, 0)
	rightQ := make([]*TreeNode, 0)
	leftQ = append(leftQ, root)
	rightQ = append(rightQ, root)

	for len(leftQ) != 0 && len(rightQ) != 0 {
		leftCur, rightCur := leftQ[0], rightQ[0]
		leftQ, rightQ = leftQ[1:], rightQ[1:]

		if leftCur == nil && rightCur == nil {
			continue
		} else if leftCur != nil && rightCur != nil && leftCur.Val == rightCur.Val {
			leftQ = append(leftQ, leftCur.Left, leftCur.Right)
			rightQ = append(rightQ, rightCur.Right, rightCur.Left)
		} else {
			return false
		}
	}

	if len(leftQ) == 0 && len(rightQ) == 0 {
		return true
	} else {
		return false
	}
}
```

## 104.二叉树的最大深度

### 题目

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 迭代 | O(n)       | O(n)       |

### 解答思路

```go
// 递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	return max(left, right) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


// 迭代
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := make([]*TreeNode,0)
	queue = append(queue,root)
	depth := 0

	for len(queue) > 0{
		length := len(queue)

		for i := 0; i < length; i++{
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil{
				queue = append(queue,node.Left)
			}
			if node.Right != nil{
				queue = append(queue,node.Right)
			}
		}
		depth++
	}
	return depth
}
```

## 107.二叉树的层次遍历II

### 题目

````
给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

例如：
给定二叉树 [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7

返回其自底向上的层次遍历为：
[
  [15,7],
  [9,20],
  [3]
]
````

### 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(n)       |
| 02   | 迭代 | O(n)       | O(n)       |

```go
// 迭代
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode,0)
	out := make([][]int,0)
	queue = append(queue, root)

	for len(queue) != 0 {
		l := len(queue)
		arr := make([]int,0)
		for i := 0; i < l; i++ {
			pop := queue[i]
			arr = append(arr, pop.Val)
			if pop.Left != nil {
				queue = append(queue, pop.Left)
			}
			if pop.Right != nil {
				queue = append(queue, pop.Right)
			}
		}
		out = append(out, arr)
		queue = queue[l:]
	}

	out2 := make([][]int, len(out))
	for i := 0; i < len(out); i++ {
		out2[len(out)-1-i] = out[i]
	}

	return out2
}

// 递归
func levelOrderBottom(root *TreeNode) [][]int {
	result := make([][]int, 0)
	level := 0
	if root == nil {
		return result
	}

	orderBottom(root, &result, level)

	left, right := 0, len(result)-1
	for left < right {
		result[left], result[right] = result[right], result[left]
		left++
		right--
	}
	return result
}

func orderBottom(root *TreeNode, result *[][]int, level int) {
	if root == nil {
		return
	}
	if len(*result) > level {
		fmt.Println(level, result, root.Val)
		(*result)[level] = append((*result)[level], root.Val)
	} else {
		*result = append(*result, []int{root.Val})
	}
	orderBottom(root.Left, result, level+1)
	orderBottom(root.Right, result, level+1)
}
```

## 108.将有序数组转换为二叉搜索树

### 题目

```
将一个按照升序排列的有序数组，转换为一棵高度平衡二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

示例:

给定有序数组: [-10,-3,0,5,9],

一个可能的答案是：[0,-3,9,-10,null,5]，它可以表示下面这个高度平衡二叉搜索树：

      0
     / \
   -3   9
   /   /
 -10  5
```

### 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 递归 | O(n)       | O(log(n))  |
| 02   | 迭代 | O(n)       | O(n)       |

```go
// 递归
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}

// 迭代
type MyTreeNode struct {
	root  *TreeNode
	start int
	end   int
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	queue := make([]MyTreeNode, 0)
	root := &TreeNode{Val: 0}
	queue = append(queue, MyTreeNode{root, 0, len(nums)})
	for len(queue) > 0 {
		myRoot := queue[0]
		queue = queue[1:]
		start := myRoot.start
		end := myRoot.end
		mid := (start + end) / 2
		curRoot := myRoot.root
		curRoot.Val = nums[mid]
		if start < mid {
			curRoot.Left = &TreeNode{Val: 0}
			queue = append(queue, MyTreeNode{curRoot.Left, start, mid})
		}
		if mid+1 < end {
			curRoot.Right = &TreeNode{Val: 0}
			queue = append(queue, MyTreeNode{curRoot.Right, mid + 1, end})
		}
	}
	return root
}
```

