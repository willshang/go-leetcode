package main

func main() {

}

// leetcode547_朋友圈
func findCircleNum(M [][]int) int {
	n := len(M)
	arr := make([]bool, n)
	res := 0
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if arr[i] == false {
			queue = append(queue, i)
			for len(queue) > 0 {
				node := queue[0]
				queue = queue[1:]
				arr[node] = true
				for j := 0; j < n; j++ {
					if M[node][j] == 1 && arr[j] == false {
						queue = append(queue, j)
					}
				}
			}
			res++
		}
	}
	return res
}
