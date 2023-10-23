package main

func main() {

}

func strWithout3a3b(a int, b int) string {
	res := make([]byte, 0)
	for a > 0 && b > 0 {
		if a > b {
			res = append(res, []byte{'a', 'a', 'b'}...)
			a = a - 2
			b--
		} else if a == b {
			res = append(res, []byte{'a', 'b'}...)
			a--
			b--
		} else {
			res = append(res, []byte{'b', 'b', 'a'}...)
			a--
			b = b - 2
		}
	}
	for a > 0 {
		res = append(res, 'a')
		a--
	}
	for b > 0 {
		res = append(res, 'b')
		b--
	}
	return string(res)
}
