package main

func main() {

}

func numberOfSubstrings(s string) int {
	res := 0
	n := len(s)
	m := make(map[int]int)
	count := 0
	var value int
	i := 0
	for j := 0; j < n; j++ {
		value = int(s[j] - 'a')
		if m[value] == 0 {
			count++
		}
		m[value]++
		for count == 3 {
			res = res + n - j
			value = int(s[i] - 'a')
			m[value]--
			i++
			if m[value] == 0 {
				count--
			}
		}
	}
	return res
}
