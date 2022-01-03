package main

import "sort"

func main() {

}

// leetcode2126_摧毁小行星
func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	temp := int64(mass)
	for i := 0; i < len(asteroids); i++ {
		if temp < int64(asteroids[i]) {
			return false
		}
		temp = temp + int64(asteroids[i])
	}
	return true
}
