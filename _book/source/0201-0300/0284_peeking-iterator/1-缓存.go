package main

func main() {

}

type Iterator struct {
}

func (this *Iterator) hasNext() bool {

}

func (this *Iterator) next() int {

}

type PeekingIterator struct {
	Iter    *Iterator
	cache   int
	isCache bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{
		Iter:    iter,
		cache:   0,
		isCache: false,
	}
}

func (this *PeekingIterator) hasNext() bool {
	return this.isCache || this.Iter.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.isCache == false {
		return this.Iter.next()
	}
	res := this.cache
	this.isCache = false
	return res
}

func (this *PeekingIterator) peek() int {
	if this.isCache == false {
		this.cache = this.Iter.next()
		this.isCache = true
	}
	return this.cache
}
