package main

func main() {

}

// leetcode1766_互质树
var res []int             // 结果
var arr [][]int           // 邻接表
var path map[int][][2]int // 路径：以每个值做区分来记录path，每个值取最新的path来比较即可

func getCoprimes(nums []int, edges [][]int) []int {
	n := len(nums)
	res = make([]int, n)
	arr = make([][]int, n)
	path = make(map[int][][2]int)
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	dfs(nums, 0, -1, 0)
	return res
}

func dfs(nums []int, cur int, prev int, level int) {
	index, depth := -1, -1
	for i := 1; i <= 50; i++ {
		if len(path[i]) > 0 {
			a, b := path[i][len(path[i])-1][0], path[i][len(path[i])-1][1] // 遍历每个值最近
			if a > depth && gcd(i, nums[cur]) == 1 {                       // 更新：更近+互质
				depth = a
				index = b
			}
		}
	}
	res[cur] = index
	for i := 0; i < len(arr[cur]); i++ {
		value := nums[cur]
		next := arr[cur][i]
		if next != prev {
			path[value] = append(path[value], [2]int{level, cur})
			dfs(nums, next, cur, level+1)
			path[value] = path[value][:len(path[value])-1]
		}
	}
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
