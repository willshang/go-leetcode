package main

func main() {

}

// leetcode2069_模拟行走机器人II
var m = map[int]string{0: "East", 1: "North", 2: "West", 3: "South"}
var dx = []int{1, 0, -1, 0}
var dy = []int{0, 1, 0, -1}

type Robot struct {
	w, h, x, y, dir, total int
}

func Constructor(width int, height int) Robot {
	return Robot{w: width, h: height, total: 2*width + 2*height - 4}
}

func (this *Robot) Step(num int) {
	num = num % this.total
	if num == 0 && this.x == 0 && this.y == 0 && this.dir == 0 {
		this.dir = 3 // 注意特判
	}
	for ; num > 0; num-- {
		newX, newY := this.x+dx[this.dir], this.y+dy[this.dir]
		if 0 <= newX && newX < this.w && 0 <= newY && newY < this.h {
			this.x = newX
			this.y = newY
		} else {
			this.dir = (this.dir + 1) % 4
			this.x = this.x + dx[this.dir]
			this.y = this.y + dy[this.dir]
		}
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.y}
}

func (this *Robot) GetDir() string {
	return m[this.dir%4]
}
