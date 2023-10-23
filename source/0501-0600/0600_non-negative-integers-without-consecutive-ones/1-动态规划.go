package main

func main() {

}

// leetcode600_不含连续1的非负整数
func findIntegers(n int) int {
	dp := make(map[int]int) // dp[i]=>高度为i，根节点为0的满二叉树，不包含连续1的路径数量（下标从1开始）
	dp[1] = 1               // 高度为1的时候，只有0等1种情况
	dp[2] = 2               // 高度为2的时候，有00、01等2种情况
	for i := 3; i <= 32; i++ {
		dp[i] = dp[i-1] + dp[i-2] // 左子树(0)+右子树的左子树(10)的数量
	}
	res := 0
	prev := 0
	for i := 32; i >= 1; i-- { // 从最高位开始遍历进行替换（下标从1开始）
		if n&(1<<(i-1)) > 0 { // 第i位为1，替换该位，前缀不变
			// 1xxx1xx => 0xxxxxx => 该位变为0
			// 1xxx1xx => 1xxx0xx => 该位变为0，前缀不变
			res = res + dp[i] // 高度为i：高度=i，根节点为0的都小于n（把i位的1替换为0）
			if prev == 1 {    // 出现连续1退出，比如后面使用1011xx的前缀就不满足题目要求了
				break
			}
			prev = 1
		} else {
			prev = 0
		}
		if i == 1 { // 如果能走到最后1位，也要算上n，说明n也满足条件(小于等于n,即n也算1次)
			res++
		}
	}
	return res
}
