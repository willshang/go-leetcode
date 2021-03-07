package main

func main() {

}

type ProductOfNumbers struct {
	arr []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{arr: make([]int, 0)}
}

func (this *ProductOfNumbers) Add(num int) {
	this.arr = append(this.arr, num)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	res := 1
	n := len(this.arr)
	for i := n - 1; i >= n-k && i >= 0; i-- {
		res = res * this.arr[i]
	}
	return res
}
