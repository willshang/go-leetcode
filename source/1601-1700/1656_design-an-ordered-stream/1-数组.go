package main

func main() {

}

// leetcode1656_设计有序流
type OrderedStream struct {
	arr []string
	ptr int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		arr: make([]string, n+1),
		ptr: 1,
	}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	res := make([]string, 0)
	this.arr[id] = value
	if this.ptr < id-1 {
		return res
	}
	for i := this.ptr; i < len(this.arr); i++ {
		if this.arr[i] == "" {
			break
		}
		res = append(res, this.arr[i])
		this.ptr++
	}
	return res
}
