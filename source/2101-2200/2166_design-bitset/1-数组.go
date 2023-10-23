package main

func main() {

}

type Bitset struct {
	arr       []int
	count     int // 1的数量
	flipCount int // 反转的次量
}

func Constructor(size int) Bitset {
	return Bitset{
		arr:       make([]int, size),
		count:     0,
		flipCount: 0,
	}
}

func (this *Bitset) Fix(idx int) {
	if (this.flipCount%2 == 0 && this.arr[idx] == 0) ||
		(this.flipCount%2 == 1 && this.arr[idx] == 1) {
		this.count++
		this.arr[idx] = 1 - this.arr[idx]
	}
}

func (this *Bitset) Unfix(idx int) {
	if (this.flipCount%2 == 0 && this.arr[idx] == 1) ||
		(this.flipCount%2 == 1 && this.arr[idx] == 0) {
		this.count--
		this.arr[idx] = 1 - this.arr[idx]
	}
}

func (this *Bitset) Flip() {
	this.flipCount++
	this.count = len(this.arr) - this.count
}

func (this *Bitset) All() bool {
	return len(this.arr) == this.count
}

func (this *Bitset) One() bool {
	return this.count > 0
}

func (this *Bitset) Count() int {
	return this.count
}

func (this *Bitset) ToString() string {
	temp := make([]byte, len(this.arr))
	for i := 0; i < len(this.arr); i++ {
		v := this.arr[i]
		if this.flipCount%2 == 1 {
			v = 1 - v
		}
		temp[i] = byte('0' + v)
	}
	return string(temp)
}
