package main

func main() {

}

type StreamRank struct {
	length int
	c      []int
}

func Constructor() StreamRank {
	return StreamRank{
		length: 50002,
		c:      make([]int, 50003),
	}
}

func (this *StreamRank) Track(x int) {
	this.upData(x+1, 1)
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	return this.getSum(x + 1)
}

func (this *StreamRank) lowBit(x int) int {
	return x & (-x)
}

// 单点修改
func (this *StreamRank) upData(i, k int) { // 在i位置加上k
	for i <= this.length {
		this.c[i] = this.c[i] + k
		i = i + this.lowBit(i) // i = i + 2^k
	}
}

// 区间查询
func (this *StreamRank) getSum(i int) int {
	res := 0
	for i > 0 {
		res = res + this.c[i]
		i = i - this.lowBit(i)
	}
	return res
}
