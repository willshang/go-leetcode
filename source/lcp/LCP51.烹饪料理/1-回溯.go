package main

func main() {

}

var res int

func perfectMenu(materials []int, cookbooks [][]int, attribute [][]int, limit int) int {
	res = -1
	arr := make([]int, 5)
	dfs(materials, cookbooks, attribute, limit, 0, 0, arr, 0)
	return res
}

func dfs(materials []int, cookbooks [][]int, attribute [][]int, limit int, sumX, sumY int, arr []int, index int) {
	flag := true
	for i := 0; i < 5; i++ {
		if arr[i] > materials[i] {
			flag = false
			break
		}
	}
	if flag == false {
		return
	}
	if sumY >= limit {
		res = max(res, sumX)
	}
	if index >= len(cookbooks) {
		return
	}
	tempA, tempB := make([]int, 5), make([]int, 5)
	copy(tempA, arr)
	copy(tempB, arr)
	for i := 0; i < 5; i++ {
		tempA[i] = tempA[i] + cookbooks[index][i]
	}
	dfs(materials, cookbooks, attribute, limit, sumX+attribute[index][0], sumY+attribute[index][1], tempA, index+1) // 加上
	dfs(materials, cookbooks, attribute, limit, sumX, sumY, tempB, index+1)                                         // 不加上
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
