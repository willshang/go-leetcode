package main

func main() {

}

// leetcode2105_给植物浇水II
func minimumRefill(plants []int, capacityA int, capacityB int) int {
	res := 0
	i, j := 0, len(plants)-1
	a, b := capacityA, capacityB
	for i <= j {
		if i == j { // 相等
			if a >= b && a < plants[i] {
				res++
			}
			if a < b && b < plants[i] {
				res++
			}
			return res
		}
		if a < plants[i] {
			res++
			a = capacityA - plants[i]
		} else {
			a = a - plants[i]
		}
		i++
		if b < plants[j] {
			res++
			b = capacityB - plants[j]
		} else {
			b = b - plants[j]
		}
		j--
	}
	return res
}
