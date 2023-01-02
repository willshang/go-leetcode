package main

import (
	"fmt"
	"math"
)

func main() {
	var r, x, y, x1, y1 int
	fmt.Scan(&r, &x, &y, &x1, &y1)
	res := getResult(r, x, y, x1, y1)
	fmt.Println(res)
}

func getResult(r, x, y, x1, y1 int) int {
	v := (x-x1)*(x-x1) + (y-y1)*(y-y1)
	dis := math.Sqrt(float64(v))
	res := int(dis / float64(2*r)) // 除以直径
	if float64(res) < dis/float64(2*r) {
		res = res + 1
	}
	return res
}
