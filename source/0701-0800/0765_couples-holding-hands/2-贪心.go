package main

func main() {

}

// leetcode765_情侣牵手
func minSwapsCouples(row []int) int {
	res := 0
	for i := 0; i < len(row); i = i + 2 {
		a, b := row[i], row[i+1]
		if b == a^1 {
			continue
		}
		res = res + 1
		for j := i + 1; j < len(row); j++ {
			if row[j] == a^1 {
				row[j], row[i+1] = row[i+1], row[j]
				break
			}
		}
	}
	return res
}
