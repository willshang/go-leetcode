package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(poorPigs(1000, 15, 59))
}

func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	count := minutesToTest/minutesToDie + 1 // 最多喝几次水
	res := math.Log(float64(buckets)) / math.Log(float64(count))
	return int(math.Ceil(res))
}
