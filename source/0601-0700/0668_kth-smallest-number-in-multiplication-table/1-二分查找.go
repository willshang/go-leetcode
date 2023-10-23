package main

func main() {

}

// leetcode668_乘法表中第k小的数
func findKthNumber(m int, n int, k int) int {
	left := 1
	right := m * n
	for left < right {
		mid := left + (right-left)/2
		total := judge(m, n, k, mid)
		if total == true { // 满足条件，继续找
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func judge(m int, n int, k int, target int) bool {
	count := 0
	for i := 1; i <= m; i++ {
		count = count + min(target/i, n) // 当前行全部满足+n，部分满足+target/i
	}
	return count >= k
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
