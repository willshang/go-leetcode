package main

func main() {

}

// leetcode355_设计推特
type Twitter struct {
	data [][2]int
	m    map[int]map[int]int
}

func Constructor() Twitter {
	return Twitter{
		data: make([][2]int, 0),
		m:    make(map[int]map[int]int),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.data = append(this.data, [2]int{userId, tweetId})
}

func (this *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0)
	for i := len(this.data) - 1; i >= 0; i-- { // 遍历发表的列表
		id, tid := this.data[i][0], this.data[i][1]
		if id == userId || this.m[userId][id] > 0 {
			res = append(res, tid)
		}
		if len(res) == 10 {
			return res
		}
	}
	return res
}

func (this *Twitter) Follow(followerId int, followeeId int) {
	if this.m[followerId] == nil {
		this.m[followerId] = make(map[int]int)
	}
	this.m[followerId][followeeId] = 1
}

func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if this.m[followerId] != nil {
		this.m[followerId][followeeId] = 0
	}
}
