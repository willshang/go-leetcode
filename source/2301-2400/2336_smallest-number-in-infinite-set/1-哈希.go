package main

func main() {

}

// leetcode2336_无限集中的最小数字
type SmallestInfiniteSet struct {
	m map[int]bool
}

func Constructor() SmallestInfiniteSet {
	return SmallestInfiniteSet{m: make(map[int]bool)}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	for i := 1; i <= 1000; i++ {
		if this.m[i] == false {
			this.m[i] = true
			return i
		}
	}
	return 0
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	delete(this.m, num)
}
