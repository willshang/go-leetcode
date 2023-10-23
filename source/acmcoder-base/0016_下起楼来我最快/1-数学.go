package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var t1, t2, t3, t4 int
	fmt.Scan(&t1, &t2, &t3, &t4)
	res := getResult(n, m, t1, t2, t3, t4)
	fmt.Println(res)
}

func getResult(n, m, t1, t2, t3, t4 int) int {
	a := t4 * (n - 1)                    // 纯爬楼梯
	dis := int(math.Abs(float64(n - m))) // 楼层差距
	b := t2*2 + t3 + dis*t1 + (n-1)*t1   // 先爬楼梯到电梯所在层，即m层，然后坐电梯=>开门x2+关门+爬楼梯+电梯运行
	if a < b {
		return a
	}
	return b
}
