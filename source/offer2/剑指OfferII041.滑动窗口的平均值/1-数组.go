package main

func main() {

}

// 剑指OfferII041.滑动窗口的平均值
type MovingAverage struct {
	sum  int
	size int
	arr  []int
}

func Constructor(size int) MovingAverage {
	return MovingAverage{
		sum:  0,
		size: size,
		arr:  make([]int, 0),
	}
}

func (this *MovingAverage) Next(val int) float64 {
	this.sum = this.sum + val
	this.arr = append(this.arr, val)
	if len(this.arr) > this.size {
		this.sum = this.sum - this.arr[0]
		this.arr = this.arr[1:]
	}
	return float64(this.sum) / float64(len(this.arr))
}
