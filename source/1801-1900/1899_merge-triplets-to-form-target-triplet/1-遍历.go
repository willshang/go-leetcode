package main

func main() {

}

func mergeTriplets(triplets [][]int, target []int) bool {
	x, y, z := target[0], target[1], target[2]
	a, b, c := 0, 0, 0
	for i := 0; i < len(triplets); i++ {
		a1, b1, c1 := triplets[i][0], triplets[i][1], triplets[i][2]
		if a1 <= x && b1 <= y && c1 <= z {
			a = max(a, a1)
			b = max(b, b1)
			c = max(c, c1)
		}
	}
	return a == x && b == y && c == z
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
