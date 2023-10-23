package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(numberOfRounds("12:04", "12:44"))
}

// leetcode1904_你完成的完整对局数
func numberOfRounds(startTime string, finishTime string) int {
	a, _ := strconv.Atoi(startTime[:2])
	b, _ := strconv.Atoi(startTime[3:])
	c, _ := strconv.Atoi(finishTime[:2])
	d, _ := strconv.Atoi(finishTime[3:])
	t1 := 60*a + b
	t2 := 60*c + d
	if t2 < t1 {
		t2 = t2 + 1440
	}
	start := int(math.Ceil(float64(t1) / float64(15)))   // 向上取整
	finish := int(math.Floor(float64(t2) / float64(15))) // 向下取整
	if finish-start >= 0 {
		return finish - start
	}
	return 0 // 00:47 ~00:48
}
