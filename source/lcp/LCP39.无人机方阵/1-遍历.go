package main

func main() {

}

// leetcode_lcp39_无人机方阵
func minimumSwitchingTimes(source [][]int, target [][]int) int {
	temp := make(map[int]int)
	n, m := len(source), len(source[0])
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			temp[source[i][j]]++
			temp[target[i][j]]--
		}
	}
	res := 0
	for _, v := range temp {
		if v > 0 {
			res = res + v
		}
	}
	return res
}
