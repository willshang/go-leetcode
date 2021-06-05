package main

func main() {

}

// leetcode1348_推文计数
type TweetCounts struct {
	m map[string][]int
}

func Constructor() TweetCounts {
	return TweetCounts{m: make(map[string][]int)}
}

func (this *TweetCounts) RecordTweet(tweetName string, time int) {
	this.m[tweetName] = append(this.m[tweetName], time)
}

func (this *TweetCounts) GetTweetCountsPerFrequency(freq string, tweetName string, startTime int, endTime int) []int {
	per := 0
	if freq == "minute" {
		per = 60
	} else if freq == "hour" {
		per = 60 * 60
	} else if freq == "day" {
		per = 60 * 60 * 24
	}
	n := (endTime-startTime)/per + 1
	res := make([]int, n)
	for i := 0; i < len(this.m[tweetName]); i++ {
		t := this.m[tweetName][i]
		if startTime <= t && t <= endTime {
			index := (t - startTime) / per
			res[index]++
		}
	}
	return res
}
