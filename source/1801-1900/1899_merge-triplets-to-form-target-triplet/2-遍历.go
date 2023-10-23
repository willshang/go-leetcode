package main

func main() {

}

// leetcode1899_合并若干三元组以形成目标三元组
func mergeTriplets(triplets [][]int, target []int) bool {
	x, y, z := target[0], target[1], target[2]
	a, b, c := false, false, false
	for i := 0; i < len(triplets); i++ {
		a1, b1, c1 := triplets[i][0], triplets[i][1], triplets[i][2]
		if a1 <= x && b1 <= y && c1 <= z {
			if a1 == x {
				a = true
			}
			if b1 == y {
				b = true
			}
			if c1 == z {
				c = true
			}
		}
	}
	return a == true && b == true && c == true
}
