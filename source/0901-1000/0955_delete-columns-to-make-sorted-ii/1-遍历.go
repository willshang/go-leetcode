package main

func main() {

}

// leetcode955_删列造序II
func minDeletionSize(strs []string) int {
	res := 0
	cur := make([]string, len(strs))
	for i := 0; i < len(strs[0]); i++ {
		temp := make([]string, len(strs))
		copy(temp, cur)
		for j := 0; j < len(temp); j++ {
			temp[j] = temp[j] + string(strs[j][i])
		}
		if judge(temp) == true { // 加上当前列后，是升序的就用，不是就删除
			cur = temp
		} else {
			res++
		}
	}
	return res
}

func judge(strs []string) bool {
	for i := 0; i < len(strs)-1; i++ {
		if strs[i] > strs[i+1] {
			return false
		}
	}
	return true
}
