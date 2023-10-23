package main

func main() {

}

// leetcode_lcp55_采集果实
func getMinimumTime(time []int, fruits [][]int, limit int) int {
	res := 0
	for i := 0; i < len(fruits); i++ {
		a, b := fruits[i][0], fruits[i][1]
		res = res + ((b-1)/limit+1)*time[a]
	}
	return res
}
