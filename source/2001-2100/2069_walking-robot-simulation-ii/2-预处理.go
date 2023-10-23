package main

func main() {

}

var m = map[int]string{0: "East", 1: "North", 2: "West", 3: "South"}

type Robot struct {
	arr    [][2]int
	dir    []int
	isMove bool
	index  int
}

func Constructor(width int, height int) Robot {
	arr := make([][2]int, 0)
	dir := make([]int, 0)
	for i := 0; i < width; i++ {
		arr = append(arr, [2]int{i, 0})
		dir = append(dir, 0)
	}
	for i := 1; i < height; i++ {
		arr = append(arr, [2]int{width - 1, i})
		dir = append(dir, 1)
	}
	for i := width - 2; i >= 0; i-- {
		arr = append(arr, [2]int{i, height - 1})
		dir = append(dir, 2)
	}
	for i := height - 2; i > 0; i-- {
		arr = append(arr, [2]int{0, i})
		dir = append(dir, 3)
	}
	dir[0] = 3 // 第0个朝南
	return Robot{arr: arr, dir: dir, isMove: false}
}

func (this *Robot) Step(num int) {
	this.isMove = true
	this.index = (this.index + num) % len(this.arr)
}

func (this *Robot) GetPos() []int {
	return []int{this.arr[this.index][0], this.arr[this.index][1]}
}

func (this *Robot) GetDir() string {
	if this.isMove == false {
		return "East"
	}
	return m[this.dir[this.index]]
}
