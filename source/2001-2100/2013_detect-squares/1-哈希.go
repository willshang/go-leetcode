package main

func main() {

}

// leetcode2013_检测正方形
type DetectSquares struct {
	m map[int]map[int]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		m: make(map[int]map[int]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	a, b := point[0], point[1]
	if this.m[a] == nil {
		this.m[a] = make(map[int]int)
	}
	this.m[a][b]++
}

func (this *DetectSquares) Count(point []int) int {
	res := 0
	x, y := point[0], point[1]
	for a := range this.m[x] { // 遍历同一列的y坐标
		if a == y {
			continue
		}
		length := abs(a - y) // 获取正方形边长
		b := x + length      // 右边
		if this.m[b] != nil && this.m[b][a] > 0 && this.m[b][y] > 0 {
			res = res + this.m[b][a]*this.m[b][y]*this.m[x][a]
		}
		b = x - length // 左边
		if this.m[b] != nil && this.m[b][a] > 0 && this.m[b][y] > 0 {
			res = res + this.m[b][a]*this.m[b][y]*this.m[x][a]
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
