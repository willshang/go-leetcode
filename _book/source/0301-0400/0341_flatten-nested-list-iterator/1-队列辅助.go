package main

func main() {

}

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool {
	return true
}
func (this NestedInteger) GetInteger() int {
	return 0
}
func (n *NestedInteger) SetInteger(value int)      {}
func (this *NestedInteger) Add(elem NestedInteger) {}
func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}

type NestedIterator struct {
	arr []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{arr: nestedList}
}

func (this *NestedIterator) Next() int {
	value := this.arr[0]
	this.arr = this.arr[1:]
	return value.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	if len(this.arr) == 0 {
		return false
	}
	if this.arr[0].IsInteger() {
		return true
	}
	this.arr = append(this.arr[0].GetList(), this.arr[1:]...)
	return this.HasNext()
}
