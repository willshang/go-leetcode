package main

func main() {

}

// leetcode1386_安排电影院座位
func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	m := make(map[int]int)
	for i := 0; i < len(reservedSeats); i++ {
		a, b := reservedSeats[i][0], reservedSeats[i][1]
		if b == 1 || b == 10 { // 1和10不影响
			continue
		}
		m[a] = m[a] | (1 << (b - 2))
	}
	left := 0b11110000
	middle := 0b11000011
	right := 0b00001111
	res := (n - len(m)) * 2
	for _, v := range m {
		// v一定有1个为1
		if v|left == left || v|middle == middle || v|right == right {
			res++
		}
	}
	return res
}
