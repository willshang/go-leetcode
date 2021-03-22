package main

func main() {

}

func maxNumberOfFamilies(n int, reservedSeats [][]int) int {
	m := make(map[int]map[int]bool)
	for i := 0; i < len(reservedSeats); i++ {
		a, b := reservedSeats[i][0], reservedSeats[i][1]
		if b == 1 || b == 10 { // 1和10不影响
			continue
		}
		if m[a] == nil {
			m[a] = make(map[int]bool)
		}
		m[a][b] = true
	}
	res := 0
	for _, v := range m {
		flag := false
		// 左边
		if v[2] == false && v[3] == false && v[4] == false && v[5] == false {
			res++
			flag = true
		}
		// 右边
		if v[6] == false && v[7] == false && v[8] == false && v[9] == false {
			res++
			flag = true
		}
		// 中间
		if flag == false &&
			v[4] == false && v[5] == false && v[6] == false && v[7] == false {
			res++
		}
	}
	res = res + 2*(n-len(m))
	return res
}
