package main

func main() {

}

// 剑指OfferII058.日程表
type MyCalendar struct {
	arr [][2]int
}

func Constructor() MyCalendar {
	return MyCalendar{arr: make([][2]int, 0)}
}

func (this *MyCalendar) Book(start int, end int) bool {
	for i := 0; i < len(this.arr); i++ {
		if this.arr[i][0] < end && start < this.arr[i][1] {
			return false
		}
	}
	this.arr = append(this.arr, [2]int{start, end})
	return true
}
