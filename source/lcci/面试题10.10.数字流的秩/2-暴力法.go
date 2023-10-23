package main

// 程序员面试金典10.10_数字流的秩
type StreamRank struct {
	m map[int]int
}

func Constructor() StreamRank {
	return StreamRank{m: make(map[int]int)}
}

func (this *StreamRank) Track(x int) {
	this.m[x]++
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	res := 0
	for k, v := range this.m {
		if k <= x {
			res = res + v
		}
	}
	return res
}
