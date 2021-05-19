package main

func main() {

}

func memLeak(memory1 int, memory2 int) []int {
	i := 1
	for {
		if i <= memory1 || i <= memory2 {
			if memory1 < memory2 {
				memory2 = memory2 - i
			} else {
				memory1 = memory1 - i
			}
		} else {
			break
		}
		i++
	}
	return []int{i, memory1, memory2}
}
