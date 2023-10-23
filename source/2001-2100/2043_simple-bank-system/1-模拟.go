package main

func main() {

}

// leetcode2043_简易银行系统
type Bank struct {
	arr []int64
	n   int
}

func Constructor(balance []int64) Bank {
	return Bank{arr: balance, n: len(balance)}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > this.n || account2 > this.n || this.arr[account1-1] < money {
		return false
	}
	this.arr[account1-1] = this.arr[account1-1] - money
	this.arr[account2-1] = this.arr[account2-1] + money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > this.n {
		return false
	}
	this.arr[account-1] = this.arr[account-1] + money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > this.n || this.arr[account-1] < money {
		return false
	}
	this.arr[account-1] = this.arr[account-1] - money
	return true
}
