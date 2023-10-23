package main

func main() {

}

// leetcode1472_设计浏览器历史记录
type BrowserHistory struct {
	arr   []string
	index int
}

func Constructor(homepage string) BrowserHistory {
	return BrowserHistory{
		arr:   []string{homepage},
		index: 0,
	}
}

func (this *BrowserHistory) Visit(url string) {
	length := len(this.arr)
	if this.index == length-1 {
		this.arr = append(this.arr, url)
	} else if this.index < length-1 {
		this.arr = this.arr[:this.index+1]
		this.arr = append(this.arr, url)
	}
	this.index++
}

func (this *BrowserHistory) Back(steps int) string {
	if steps > this.index {
		this.index = 0
		return this.arr[0]
	}
	this.index = this.index - steps
	return this.arr[this.index]
}

func (this *BrowserHistory) Forward(steps int) string {
	length := len(this.arr)
	if this.index == length-1 {
	} else if this.index+steps > length-1 {
		this.index = length - 1
	} else {
		this.index = this.index + steps
	}
	return this.arr[this.index]
}
