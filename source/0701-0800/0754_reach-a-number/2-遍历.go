package main

func main() {

}

// leetcode754_到达终点数字
func reachNumber(target int) int {
	n := abs(target) // 负数转为正数，负数正数本质上都一样
	k := 1
	// S=1+...+k >= target
	// 差值：diff = S-target
	// diff为偶数：可以找到1~k之间的一个组合之和为diff/2
	// diff为奇数：需要考虑S=1~k+1或者S=1~K+2的情况，使得新diff为偶数
	// =>求S>=target 并且使用diff为偶数
	for {
		sum := k * (k + 1) / 2
		if sum >= n && (sum-n)%2 == 0 {
			return k
		}
		k++
	}
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
