package main

func main() {

}

func validSquare(p1 []int, p2 []int, p3 []int, p4 []int) bool {
	m := make(map[int]int)
	m[getDis(p1, p2)]++
	m[getDis(p1, p3)]++
	m[getDis(p1, p4)]++
	m[getDis(p2, p3)]++
	m[getDis(p2, p4)]++
	m[getDis(p3, p4)]++
	a, b := 0, 0
	for k, v := range m {
		if v == 2 {
			a = k
		} else if v == 4 {
			b = k
		} else {
			return false
		}
	}
	return len(m) == 2 && a == 2*b
}

func getDis(a, b []int) int {
	return (a[0]-b[0])*(a[0]-b[0]) + (a[1]-b[1])*(a[1]-b[1])
}
