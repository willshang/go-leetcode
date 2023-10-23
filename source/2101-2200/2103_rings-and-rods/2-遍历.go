package main

func main() {

}

func countPoints(rings string) int {
	m := make(map[byte]map[byte]bool)
	for i := 0; i < 10; i++ {
		m[byte(i+'0')] = make(map[byte]bool)
	}
	for i := 0; i < len(rings); i = i + 2 {
		m[rings[i+1]][rings[i]] = true
	}
	res := 0
	for _, v := range m {
		if len(v) == 3 {
			res++
		}
	}
	return res
}
