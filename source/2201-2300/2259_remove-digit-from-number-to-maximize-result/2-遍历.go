package main

func main() {

}

// leetcode2259_移除指定数字得到的最大结果
func removeDigit(number string, digit byte) string {
	res := ""
	for i := 0; i < len(number); i++ {
		if number[i] == digit {
			v := number[:i] + number[i+1:]
			if v > res {
				res = v
			}
		}
	}
	return res
}
