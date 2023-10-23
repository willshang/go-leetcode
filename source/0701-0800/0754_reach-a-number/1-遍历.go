package main

func main() {

}

func reachNumber(target int) int {
	n := abs(target) // 负数转为正数，负数正数本质上都一样
	k := 0
	// S=1+...+k >= target
	// 差值：diff = S-target
	// diff为偶数：可以找到1~k之间的一个组合之和为diff/2
	// diff为奇数：需要考虑S=1~k+1或者S=1~K+2的情况，使得新diff为偶数
	for n > 0 {
		k = k + 1
		n = n - k
	}
	if n%2 == 0 {
		return k
	}
	return k + 1 + k%2
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
