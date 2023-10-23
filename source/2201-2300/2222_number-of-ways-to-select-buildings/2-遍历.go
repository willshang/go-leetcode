package main

func main() {

}

// leetcode2222_选择建筑的方案数
func numberOfWays(s string) int64 {
	return getCount(s, "101") + getCount(s, "010")
}

func getCount(s string, str string) int64 {
	var a, b, c int64
	for i := 0; i < len(s); i++ {
		if s[i] == str[0] { // 以101为例：计算1的个数
			a++
		}
		if s[i] == str[1] { // 以101为例：计算01的个数
			b = b + a
		}
		if s[i] == str[2] { // 以101为例：计算101的个数
			c = c + b
		}
	}
	return c
}
