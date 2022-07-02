package main

func main() {

}

// leetcode2317_操作后的最大异或和
func maximumXOR(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		// 最大逐位异或和：在某位上有奇数个1，能达到最大
		// 更新操作：and操作能把部分位从1置为0，但是不能把0修改为1
		// 求解方案：使用更新操作把相同位置保留1个1，其它置为0 => 转为求位或
		// => 有多少位存在1
		res = res | nums[i] // XOR位或，只要该位为1，结果就是1
	}
	return res
}
