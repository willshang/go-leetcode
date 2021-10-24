package main

func main() {

}

// leetcode1361_验证二叉树
func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			arr[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			arr[rightChild[i]]++
		}
	}
	root := -1 // 寻找一个根节点
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			root = i
			break
		}
	}
	if root == -1 {
		return false
	}
	visited := make(map[int]bool)
	visited[root] = true
	queue := make([]int, 0)
	queue = append(queue, root)
	for len(queue) > 0 { // 层序遍历
		node := queue[0]
		queue = queue[1:]
		a, b := leftChild[node], rightChild[node]
		if a != -1 {
			if visited[a] == true {
				return false
			}
			visited[a] = true
			queue = append(queue, a)
		}
		if b != -1 {
			if visited[b] == true {
				return false
			}
			visited[b] = true
			queue = append(queue, b)
		}
	}
	return len(visited) == n
}
