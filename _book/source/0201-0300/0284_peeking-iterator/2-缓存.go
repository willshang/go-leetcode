package main

func main() {

}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {

}

func (this *Iterator) next() int {

}

// leetcode284_顶端迭代器
type PeekingIterator struct {
	Iter  *Iterator
	cache *int
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		Iter:  iter,
		cache: nil,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.cache != nil || this.Iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.cache != nil {
		res := *this.cache
		this.cache = nil
		return res
	}
	return this.Iter.next()
}

func (this *PeekingIterator) peek() int {
	if this.cache == nil {
		value := this.Iter.next()
		this.cache = &value
	}
	return *this.cache
}
