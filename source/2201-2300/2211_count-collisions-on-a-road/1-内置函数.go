package main

import "strings"

func main() {

}

func countCollisions(directions string) int {
	s := strings.TrimLeft(directions, "L") // 左边向左不会发生碰撞
	s = strings.TrimRight(s, "R")          // 右边向右不会发生碰撞
	return len(s) - strings.Count(s, "S")  // 剩下的车都会发生碰撞
}
