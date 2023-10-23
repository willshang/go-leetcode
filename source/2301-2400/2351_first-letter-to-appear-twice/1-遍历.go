package main

func main() {

}

func repeatedCharacter(s string) byte {
	m := make(map[byte]bool)
	for _, v := range s {
		if m[byte(v)] == true {
			return byte(v)
		}
		m[byte(v)] = true
	}
	return 0
}
