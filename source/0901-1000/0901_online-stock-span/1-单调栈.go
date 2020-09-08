package main

func main() {

}

// leetcode901_股票价格跨度
type StockSpanner struct {
	prices []int
	count  []int
}

func Constructor() StockSpanner {
	return StockSpanner{
		prices: make([]int, 0),
		count:  make([]int, 0),
	}
}

func (this *StockSpanner) Next(price int) int {
	res := 1
	for len(this.prices) > 0 && this.prices[len(this.prices)-1] <= price {
		this.prices = this.prices[:len(this.prices)-1]
		temp := this.count[len(this.count)-1]
		this.count = this.count[:len(this.count)-1]
		res = res + temp
	}
	this.prices = append(this.prices, price)
	this.count = append(this.count, res)
	return res
}
