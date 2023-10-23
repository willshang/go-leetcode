package main

func main() {

}

// leetcode2029_石子游戏IX
func stoneGameIX(stones []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(stones); i++ {
		m[stones[i]%3]++
	}
	a, b, c := m[0], m[1], m[2] // 0,1,2的个数
	// 以1开头：1、(1、2)、1...
	// 以2开头：2、(2、1)、2...
	if a%2 == 0 { // 0为偶数个，可以抵消
		// 获胜策略：选择较少的1或者2
		// 1、1=2 => Alice赢（以1开始或者2开始都赢）
		// 不等的时候
		// 2.1、1>2 => Alice以2开头=> 2、2、1...(2、1).. 最后Bob没有2可选，只能选择1，Alice就赢了
		// 2.2 2>1 => Alice以1开头=> 1、1、2...(1、2).. 最后Bob没有1可选，只能选择2，Alice就赢了
		return b > 0 && c > 0
	}
	if a%2 == 1 {
		// 奇数个0
		// 获胜策略：选取较多的1或者2
		// 需要差值大于2个Alice才赢
		// 1>2 => Alice以1开头 => 1、(1、0)、2、1、1 / 1、(0、1)、2、1、1
		// 2>1 => Alice以2开头 => 2、(2、0)、1、2、2 / 2、(0、2)、1、2、2
		return b-c > 2 || c-b > 2
	}
	return false
}
