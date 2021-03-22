package main

func main() {

}

// leetcode1629_按键持续时间最长的键
func slowestKey(releaseTimes []int, keysPressed string) byte {
	res := 0
	maxValue := releaseTimes[0]
	for i := 1; i < len(releaseTimes); i++ {
		if releaseTimes[i]-releaseTimes[i-1] > maxValue {
			maxValue = releaseTimes[i] - releaseTimes[i-1]
			res = i
		} else if releaseTimes[i]-releaseTimes[i-1] == maxValue &&
			keysPressed[i] > keysPressed[res] {
			res = i
		}
	}
	return keysPressed[res]
}
