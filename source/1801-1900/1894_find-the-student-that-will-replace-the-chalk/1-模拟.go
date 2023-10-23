package main

func main() {

}

// leetcode1894_找到需要补充粉笔的学生编号
func chalkReplacer(chalk []int, k int) int {
	res := 0
	n := len(chalk)
	sum := 0
	for i := 0; i < n; i++ {
		sum = sum + chalk[i]
	}
	k = k % sum // 求余
	for i := 0; i < n; i++ {
		if k < chalk[i] {
			return i
		}
		k = k - chalk[i]
	}
	return res
}
