package main

func main() {

}

func minSwaps(s string) int {
	count := 0
	minValue := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '[' {
			count++
		} else {
			count--
			minValue = min(minValue, count)
		}
	}
	return (abs(minValue) + 1) / 2
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
