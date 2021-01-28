package main

func main() {

}

// leetcode423_从英文中重建数字
func originalDigits(s string) string {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	// 利用字母唯一性
	arr := [10]int{}
	arr[0] = m['z']
	arr[2] = m['w']
	arr[4] = m['u']
	arr[6] = m['x']
	arr[8] = m['g']
	arr[3] = m['t'] - arr[2] - arr[8]
	arr[1] = m['o'] - arr[0] - arr[2] - arr[4]
	arr[7] = m['s'] - arr[6]
	arr[5] = m['v'] - arr[7]
	arr[9] = m['i'] - arr[5] - arr[6] - arr[8]
	res := make([]byte, 0)
	for i := 0; i < 10; i++ {
		for j := 0; j < arr[i]; j++ {
			res = append(res, byte(i+'0'))
		}
	}
	return string(res)
}
