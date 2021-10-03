package main

func main() {

}

// 剑指OfferII037.小行星碰撞
func asteroidCollision(asteroids []int) []int {
	res := make([]int, 0)
	for i := 0; i < len(asteroids); i++ {
		value := asteroids[i]
		for value < 0 && len(res) > 0 && res[len(res)-1] > 0 {
			sum := value + res[len(res)-1]
			if sum >= 0 {
				value = 0
			}
			if sum <= 0 {
				res = res[:len(res)-1]
			}
		}
		if value != 0 {
			res = append(res, value)
		}
	}
	return res
}
