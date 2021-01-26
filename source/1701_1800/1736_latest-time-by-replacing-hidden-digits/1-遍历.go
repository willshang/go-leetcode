package main

func main() {

}

// leetcode1736_替换隐藏数字得到的最晚时间
func maximumTime(time string) string {
	res := []byte(time)
	for i := 0; i < len(res); i++ {
		if res[i] == '?' {
			if i == 0 && (('0' <= time[1] && time[1] <= '3') || time[1] == '?') {
				res[i] = '2'
			} else {
				res[i] = '1'
			}
			if i == 1 {
				if time[0] == '2' || res[0] == '2' {
					res[i] = '3'
				} else {
					res[i] = '9'
				}
			}
			if i == 3 {
				res[i] = '5'
			}
			if i == 4 {
				res[i] = '9'
			}
		}
	}
	return string(res)
}
