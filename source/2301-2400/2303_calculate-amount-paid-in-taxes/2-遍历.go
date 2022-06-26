package main

func main() {

}

// leetcode2303_计算应缴税款总额
func calculateTax(brackets [][]int, income int) float64 {
	res := 0
	prev := 0
	for i := 0; i < len(brackets); i++ {
		a, b := brackets[i][0], brackets[i][1]
		if income < prev {
			break
		}
		res = res + (min(income, a)-prev)*b
		prev = a
	}
	return float64(res) / 100
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
