package main

func main() {

}

// leetcode2079_给植物浇水
func wateringPlants(plants []int, capacity int) int {
	res := 0
	n := len(plants)
	value := capacity
	for i := 0; i < n; i++ {
		if plants[i] <= value {
			res++ // 走1步
			value = value - plants[i]
		} else {
			res = res + i + (i + 1) // 回去i，返回i+1
			value = capacity - plants[i]
		}
	}
	return res
}
