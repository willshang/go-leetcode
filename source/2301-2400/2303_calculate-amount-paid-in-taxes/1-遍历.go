package main

func main() {

}

func calculateTax(brackets [][]int, income int) float64 {
	res := 0
	prev := 0
	for i := 0; i < len(brackets); i++ {
		a, b := brackets[i][0], brackets[i][1]
		if income <= a {
			res = res + (income-prev)*b
			break
		}
		res = res + (a-prev)*b
		prev = a
	}
	return float64(res) / 100
}
