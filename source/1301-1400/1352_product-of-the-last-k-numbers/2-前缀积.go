package main

func main() {

}

// leetcode1352_最后K个数的乘积
type ProductOfNumbers struct {
	arr []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{arr: []int{1}}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.arr = []int{1}
		return
	}
	num = num * this.arr[len(this.arr)-1]
	this.arr = append(this.arr, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if k > len(this.arr)-1 {
		return 0
	}
	return this.arr[len(this.arr)-1] / this.arr[len(this.arr)-1-k]
}
