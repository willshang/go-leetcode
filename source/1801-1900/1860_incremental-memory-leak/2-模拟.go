package main

func main() {

}

// leetcode1860_增长的内存泄露
func memLeak(memory1 int, memory2 int) []int {
	i := 1
	for ; i <= memory1 || i <= memory2; i++ {
		if memory1 < memory2 {
			memory2 = memory2 - i
		} else {
			memory1 = memory1 - i
		}
	}
	return []int{i, memory1, memory2}
}
