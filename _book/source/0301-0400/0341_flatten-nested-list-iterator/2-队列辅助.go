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

// leetcode341_扁平化嵌套列表迭代器
type NestedIterator struct {
	arr []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	arr := getList(nestedList)
	return &NestedIterator{arr: arr}
}

func getList(nestedList []*NestedInteger) []int {
	res := make([]int, 0)
	for i := 0; i < len(nestedList); i++ {
		if nestedList[i].IsInteger() == true {
			res = append(res, nestedList[i].GetInteger())
		} else {
			res = append(res, getList(nestedList[i].GetList())...)
		}
	}
	return res
}
func (this *NestedIterator) Next() int {
	value := this.arr[0]
	this.arr = this.arr[1:]
	return value
}

func (this *NestedIterator) HasNext() bool {
	return len(this.arr) > 0
}
