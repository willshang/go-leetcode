package main

func main() {

}

// leetcode1964_找出到每个位置为止最长的有效障碍赛跑路线
func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	n := len(obstacles)
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = 1
	}
	dp := make([]int, 0)
	dp = append(dp, obstacles[0])
	for i := 1; i < n; i++ {
		if dp[len(dp)-1] <= obstacles[i] {
			dp = append(dp, obstacles[i])
			res[i] = len(dp)
		} else {
			left, right := 0, len(dp)-1
			index := 0
			for left <= right {
				mid := left + (right-left)/2
				if dp[mid] <= obstacles[i] {
					left = mid + 1
				} else {
					index = mid
					right = mid - 1
				}
			}
			dp[index] = obstacles[i] // 替换为当前元素
			res[i] = index + 1
		}
	}
	return res
}
