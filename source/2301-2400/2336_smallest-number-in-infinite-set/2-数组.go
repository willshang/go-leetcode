package main

func main() {

}

type SmallestInfiniteSet struct {
	arr []bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{arr: make([]bool, 1001)}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; i <= 1000; i++ {
		if this.arr[i] == false {
			this.arr[i] = true
			return i
		}
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	this.arr[num] = false
}
