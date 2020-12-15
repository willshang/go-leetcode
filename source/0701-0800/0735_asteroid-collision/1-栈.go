package main

func main() {

}

// leetcode735_行星碰撞
func asteroidCollision(asteroids []int) []int {
	left := make([]int, 0)
	right := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		if asteroids[i] > 0 {
			right = append(right, asteroids[i])
		} else {
			if len(right) > 0 {
				for {
					if len(right) == 0 {
						left = append(left, asteroids[i])
						break
					}
					sum := asteroids[i] + right[len(right)-1]
					if sum == 0 {
						right = right[:len(right)-1]
						break
					} else if sum > 0 {
						break
					} else {
						right = right[:len(right)-1]
					}
				}
			} else {
				left = append(left, asteroids[i])
			}
		}

	}
	return append(left, right...)
}
