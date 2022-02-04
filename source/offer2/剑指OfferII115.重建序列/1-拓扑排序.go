package main

func main() {

}

// 剑指OfferII115.重建序列
func sequenceReconstruction(org []int, seqs [][]int) bool {
	n := len(org)
	degree := make(map[int]int) // 入度
	arr := make([][]bool, n+1)  // 邻接矩阵
	for i := 0; i < n+1; i++ {
		arr[i] = make([]bool, n+1)
	}
	for i := 0; i < len(seqs); i++ {
		for j := 0; j < len(seqs[i]); j++ {
			if seqs[i][j] < 1 || seqs[i][j] > n { // 范围不对
				return false
			}
			if _, ok := degree[seqs[i][j]]; ok == false {
				degree[seqs[i][j]] = 0 // 入度设置为0
			}
			if 0 < j {
				if arr[seqs[i][j-1]][seqs[i][j]] == false {
					arr[seqs[i][j-1]][seqs[i][j]] = true // a=>b：seqs[i][j-1] => seqs[i][j]
					degree[seqs[i][j]]++
				}
			}
		}
	}
	if len(degree) != n { // 数量不对
		return false
	}
	queue := make([]int, 0) // 拓扑排序：入度=0进队列
	for i := 1; i <= n; i++ {
		if v, ok := degree[i]; ok == true && v == 0 {
			queue = append(queue, i)
		}
	}
	index := 0 // 依次对比数据
	for len(queue) > 0 {
		length := len(queue)
		if length > 1 { // 多个入度=0不能唯一重建
			return false
		}
		if org[index] != queue[0] { // 序列不对
			return false
		}
		for i := 0; i < len(arr[queue[0]]); i++ {
			if arr[queue[0]][i] == true {
				degree[i]--
				if degree[i] == 0 {
					queue = append(queue, i)
				}
			}
		}
		queue = queue[1:]
		index++
	}
	return index == n
}
