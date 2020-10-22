package main

import "fmt"

func main() {
	fmt.Println(canMeasureWater(3, 5, 4))
}

func canMeasureWater(x int, y int, z int) bool {
	if x+y < z {
		return false
	}
	queue := make([][2]int, 0)
	queue = append(queue, [2]int{0, 0})
	m := make(map[[2]int]bool)
	for len(queue) > 0 {
		a, b := queue[0][0], queue[0][1]
		queue = queue[1:]
		if m[[2]int{a, b}] == true {
			continue
		}
		m[[2]int{a, b}] = true
		if a == z || b == z || a+b == z {
			return true
		}
		// +x
		c, d := x, b
		queue = append(queue, [2]int{c, d})
		// +y
		c, d = a, y
		queue = append(queue, [2]int{c, d})
		// -x
		c, d = 0, b
		queue = append(queue, [2]int{c, d})
		// -y
		c, d = a, 0
		queue = append(queue, [2]int{c, d})
		// x->y
		if a > y-b {
			c, d = a+b-y, y
			queue = append(queue, [2]int{c, d})
		} else {
			c, d = 0, a+b
			queue = append(queue, [2]int{c, d})
		}
		// y->x
		if b > x-a {
			c, d = x, a+b-x
			queue = append(queue, [2]int{c, d})
		} else {
			c, d = a+b, 0
			queue = append(queue, [2]int{c, d})
		}
	}
	return false
}
