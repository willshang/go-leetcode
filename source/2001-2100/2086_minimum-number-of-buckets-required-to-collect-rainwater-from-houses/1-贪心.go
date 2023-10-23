package main

func main() {

}

// leetcode2086_从房屋收集雨水需要的最少水桶数
func minimumBuckets(street string) int {
	res := 0
	n := len(street)
	for i := 0; i < n; i++ {
		if street[i] == 'H' {
			if i+1 < n && street[i+1] == '.' { // (H.A)BC的情况，往后数3(+2再+1)位到B，+1
				res++
				i = i + 2
			} else if i >= 1 && street[i-1] == '.' { // (.HH)ABC的情况，+1
				res++
			} else {
				return -1
			}
		}
	}
	return res
}
