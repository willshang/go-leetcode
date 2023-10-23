package main

func main() {

}

// leetcode1238_循环码排列
func circularPermutation(n int, start int) []int {
	total := 1 << n
	res := make([]int, total)
	for i := 0; i < total; i++ {
		// 计算格雷码：leetcode89.格雷编码
		res[i] = i ^ (i >> 1) ^ start // 在计算格雷码的基础上，再与start异或
	}
	return res
}
