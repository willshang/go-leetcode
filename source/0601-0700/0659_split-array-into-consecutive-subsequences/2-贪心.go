package main

func main() {

}

// leetcode659_分割数组为连续子序列
func isPossible(nums []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		m[nums[i]]++
	}
	count := make(map[int]int) // 以某个数字结尾的连续子序列的个数
	for i := 0; i < len(nums); i++ {
		v := nums[i]
		if m[v] == 0 {
			continue
		}
		if count[v-1] > 0 { // 添加到前一个数之后
			m[v]--
			count[v-1]--
			count[v]++
		} else if m[v+1] > 0 && m[v+2] > 0 { // 没有，生成1个新的
			m[v]--
			m[v+1]--
			m[v+2]--
			count[v+2]++
		} else {
			return false
		}
	}
	return true
}
