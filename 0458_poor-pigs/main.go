package main

import "fmt"

func main() {
	fmt.Println(poorPigs(1000,15,59))
}
func poorPigs(buckets int, minutesToDie int, minutesToTest int) int {
	count := minutesToTest / minutesToDie + 1
	ret := 0
	power := 1

	for power < buckets{
		power = power * count
		ret++
	}

	return ret
}