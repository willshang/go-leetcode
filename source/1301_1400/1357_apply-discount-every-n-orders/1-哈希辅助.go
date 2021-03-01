package main

func main() {

}

// leetcode1357_每隔n个顾客打折
type Cashier struct {
	n        int
	discount int
	count    int
	m        map[int]int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	res := Cashier{
		n:        n,
		discount: discount,
		count:    0,
		m:        make(map[int]int),
	}
	for i := 0; i < len(products); i++ {
		res.m[products[i]] = prices[i]
	}
	return res
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	var res int
	for i := 0; i < len(product); i++ {
		res = res + amount[i]*this.m[product[i]]
	}
	this.count++
	if this.count%this.n == 0 {
		this.count = 0
		return float64(res*(100-this.discount)) / 100.0
	}
	return float64(res)
}
