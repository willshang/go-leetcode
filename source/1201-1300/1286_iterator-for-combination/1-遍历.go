package main

func main() {

}

type CombinationIterator struct {
	flag bool
	s    string
	arr  []int
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	arr := make([]int, combinationLength)
	for i := 0; i < combinationLength; i++ {
		arr[i] = i
	}
	return CombinationIterator{
		flag: false,
		s:    characters,
		arr:  arr,
	}
}

func (this *CombinationIterator) Next() string {
	res := ""
	for i := 0; i < len(this.arr); i++ {
		res = res + string(this.s[this.arr[i]])
	}
	index := -1
	for i := len(this.arr) - 1; i >= 0; i-- {
		// 正常情况下：以abcdef 3 为例子
		// 0 1 5 => abf  下一个 acd
		// 6-3+2 = 5
		// 6-3+1 != 1 => index = 1
		// 然后: arr[index]+1，后面递增
		target := len(this.s) - len(this.arr) + i
		if this.arr[i] != target { // 从右到左边：找到
			index = i
			break
		}
	}
	if index == -1 { // 没有更大的
		this.flag = true
	} else {
		this.arr[index]++
		for i := index + 1; i < len(this.arr); i++ {
			this.arr[i] = this.arr[i-1] + 1
		}
	}
	return res
}

func (this *CombinationIterator) HasNext() bool {
	return this.flag == false
}
