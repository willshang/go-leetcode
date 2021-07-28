package main

import "fmt"

func main() {
	re := Constructor([]int{3, 8, 0, 9, 2, 5})
	fmt.Println(re)
	fmt.Println(re.Next(2), re.cur)
	fmt.Println(re.Next(1), re.cur)
	fmt.Println(re.Next(1))
}

type RLEIterator struct {
	arr    []int
	total  int
	values []int
	cur    int
}

func Constructor(encoding []int) RLEIterator {
	total := 0
	arr := make([]int, 0) // 前缀和
	values := make([]int, 0)
	for i := 0; i < len(encoding); i = i + 2 {
		if encoding[i] == 0 {
			continue
		}
		total = total + encoding[i]
		arr = append(arr, total)
		values = append(values, encoding[i+1])
	}
	return RLEIterator{
		arr:    arr,
		total:  total,
		cur:    0,
		values: values,
	}
}

func (this *RLEIterator) Next(n int) int {
	target := this.cur + n
	this.cur = target
	if target > this.total {
		return -1
	}
	left, right := 0, len(this.arr)
	for left < right {
		mid := left + (right-left)/2
		if this.arr[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return this.values[left]
}
