package main

func main() {

}

// leetcode1797_设计一个验证系统
type AuthenticationManager struct {
	m          map[string]int
	timeToLive int
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		m:          make(map[string]int),
		timeToLive: timeToLive,
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.m[tokenId] = currentTime + this.timeToLive
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	if v, ok := this.m[tokenId]; ok {
		if v <= currentTime {
			delete(this.m, tokenId)
		} else {
			this.m[tokenId] = currentTime + this.timeToLive
		}
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	count := 0
	for _, v := range this.m {
		if v > currentTime {
			count++
		}
	}
	return count
}
