package main

func main() {

}

func numFriendRequests(ages []int) int {
	res := 0
	arr := [121]int{}
	for i := 0; i < len(ages); i++ {
		arr[ages[i]]++
	}
	for a := 0; a <= 120; a++ {
		countA := arr[a]
		for b := 0; b <= 120; b++ {
			countB := arr[b]
			if a/2+7 >= b {
				continue
			}
			if a < b {
				continue
			}
			if a < 100 && 100 < b {
				continue
			}
			res = res + countA*countB
			if a == b {
				res = res - countA
			}
		}
	}
	return res
}
