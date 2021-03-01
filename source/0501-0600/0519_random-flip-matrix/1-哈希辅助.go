package main

import "math/rand"

func main() {

}

// leetcode519_随机翻转矩阵
type Solution struct {
	m     map[int]bool
	rows  int
	cols  int
	total int
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{
		m:     make(map[int]bool),
		rows:  n_rows,
		cols:  n_cols,
		total: 0,
	}
}

func (this *Solution) Flip() []int {
	if this.total >= this.rows*this.cols {
		return nil
	}
	for {
		index := rand.Intn(this.rows * this.cols)
		if this.m[index] == true {
			continue
		}
		a, b := index/this.cols, index%this.cols
		this.m[index] = true
		this.total++
		return []int{a, b}
	}
}

func (this *Solution) Reset() {
	this.m = make(map[int]bool)
	this.total = 0
}
