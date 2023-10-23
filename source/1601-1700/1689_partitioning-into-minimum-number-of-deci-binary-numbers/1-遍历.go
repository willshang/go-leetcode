package main

func main() {

}

// leetcode1689_十-二进制数的最少数目
func minPartitions(n string) int {
	res := 0
	for i := 0; i < len(n); i++ {
		value := int(n[i] - '0')
		if value > res {
			res = value
		}
	}
	return res
}
