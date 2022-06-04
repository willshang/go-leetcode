package main

func main() {

}

// leetcode672_灯泡开关Ⅱ
func flipLights(n int, presses int) int {
	// 找规律：
	// 同一按钮操作2次，结果不变
	// 操作状态有16种；m > 4 后，状态在m=3或者m=4之间切换
	// m = 0： 0000
	// m = 1: 1000、0100、0010、0001
	// m = 2:
	//	还原：0000
	//  新增：1100、1010、1001，0110、0101、0011
	// m = 3:
	// 还原：1000、0100、0010、0001
	// 新增: 1110、0111、1011、1101
	// m = 4:
	// 还原：0000、1100、1010、1001，0110、0101、0011
	// 新增：1111
	if presses == 0 {
		return 1
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		if presses == 1 {
			return 3
		}
		return 4
	}
	if presses == 1 {
		return 4
	} else if presses == 2 {
		return 7
	}
	return 8
}
