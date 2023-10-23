package main

func main() {

}

// leetcode731_我的日程安排表II
type MyCalendarTwo struct {
	first  [][2]int
	second [][2]int
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{}
}

func (this *MyCalendarTwo) Book(start int, end int) bool {
	for i := 0; i < len(this.second); i++ {
		a, b := this.second[i][0], this.second[i][1]
		if start < b && end > a { // 跟第二个重复
			return false
		}
	}
	for i := 0; i < len(this.first); i++ {
		a, b := this.first[i][0], this.first[i][1]
		if start < b && end > a {
			// 插入重叠的部分
			this.second = append(this.second, [2]int{max(start, a), min(end, b)})
		}
	}
	this.first = append(this.first, [2]int{start, end})
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
