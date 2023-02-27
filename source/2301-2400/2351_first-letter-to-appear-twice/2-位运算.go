package main

func main() {

}

func repeatedCharacter(s string) byte {
	mask := 0
	for _, v := range s {
		target := 1 << (v - 'a')
		if mask&target != 0 {
			return byte(v)
		}
		mask = mask | target
	}
	return 0
}
