package main

func main() {

}

// leetcode1936_新增的最少台阶数
func addRungs(rungs []int, dist int) int {
	res := 0
	cur := 0
	for i := 0; i < len(rungs); i++ {
		res = res + (rungs[i]-cur-1)/dist
		cur = rungs[i]
	}
	return res
}
