package main

func main() {

}

// leetcode2139_得到目标值的最少行动次数
func minMoves(target int, maxDoubles int) int {
	res := 0
	// 方向操作：从大到小，优先减半
	for maxDoubles > 0 && target != 1 {
		res++
		if target%2 == 1 { // 奇数减去1
			target--
		} else { // 偶数/2
			maxDoubles--
			target = target / 2
		}
	}
	res = res + (target - 1)
	return res
}
