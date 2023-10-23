package main

func main() {

}

func timeRequiredToBuy(tickets []int, k int) int {
	res := 1
	for {
		for i := 0; i < len(tickets); i++ {
			if tickets[i] == 0 {
				continue
			}
			tickets[i]--
			if tickets[i] == 0 && i == k {
				return res
			}
			res++
		}
	}
}
