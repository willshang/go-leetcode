package main

import "math"

func main() {

}

// leetcode1344_时钟指针的夹角
func angleClock(hour int, minutes int) float64 {
	m := float64(minutes) * 6
	h := float64(hour)*30 + float64(minutes)*0.5
	res := math.Abs(m - h)
	if res > 180 {
		return 360 - res
	}
	return res
}
