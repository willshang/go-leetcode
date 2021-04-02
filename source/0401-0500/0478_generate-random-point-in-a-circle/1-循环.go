package main

import "math/rand"

func main() {

}

// leetcode478_在圆内随机生成点
type Solution struct {
	x      float64
	y      float64
	radius float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{
		x:      x_center,
		y:      y_center,
		radius: radius,
	}
}

func (this *Solution) RandPoint() []float64 {
	for {
		a := this.x - this.radius + 2*this.radius*rand.Float64()
		b := this.y - this.radius + 2*this.radius*rand.Float64()
		if (a-this.x)*(a-this.x)+(b-this.y)*(b-this.y) < this.radius*this.radius {
			return []float64{a, b}
		}
	}
}
