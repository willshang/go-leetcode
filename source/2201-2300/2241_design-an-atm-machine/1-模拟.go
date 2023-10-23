package main

func main() {

}

// leetcode2241_设计一个ATM机器
var Array = [5]int{20, 50, 100, 200, 500}

type ATM struct {
	arr [5]int
}

func Constructor() ATM {
	return ATM{arr: [5]int{}}
}

func (this *ATM) Deposit(banknotesCount []int) {
	for i := 0; i < len(banknotesCount); i++ {
		this.arr[i] = this.arr[i] + banknotesCount[i]
	}
}

func (this *ATM) Withdraw(amount int) []int {
	res := make([]int, 5)
	for i := 4; i >= 0; i-- {
		res[i] = min(this.arr[i], amount/Array[i])
		amount = amount - Array[i]*res[i]
	}
	if amount > 0 {
		return []int{-1}
	}
	for i := 0; i < 5; i++ {
		this.arr[i] = this.arr[i] - res[i]
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
