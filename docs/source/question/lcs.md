# LCS

## LCS01.下载插件(3)

- 题目

```
小扣打算给自己的 VS code 安装使用插件，初始状态下带宽每分钟可以完成 1 个插件的下载。假定每分钟选择以下两种策略之一:
使用当前带宽下载插件
将带宽加倍（下载插件数量随之加倍）
请返回小扣完成下载 n 个插件最少需要多少分钟。
注意：实际的下载的插件数量可以超过 n 个
示例 1：输入：n = 2 输出：2
解释：以下两个方案，都能实现 2 分钟内下载 2 个插件
方案一：第一分钟带宽加倍，带宽可每分钟下载 2 个插件；第二分钟下载 2 个插件
方案二：第一分钟下载 1 个插件，第二分钟下载 1 个插件
示例 2：输入：n = 4 输出：3
解释：最少需要 3 分钟可完成 4 个插件的下载，以下是其中一种方案:
第一分钟带宽加倍，带宽可每分钟下载 2 个插件;
第二分钟下载 2 个插件;
第三分钟下载 2 个插件。
提示：1 <= n <= 10^5
```

- 解题思路

| No. | 思路   | 时间复杂度     | 空间复杂度 |
|:----|:-----|:----------|:------|
| 01  | 遍历   | O(log(n)) | O(1)  |
| 02  | 内置函数 | O(log(n)) | O(1)  |
| 03  | 内置函数 | O(log(n)) | O(1)  |

```go
func leastMinutes(n int) int {
	res := 0
	for i := 1; i < n; i = i * 2 {
		res++
	}
	return res + 1
}

# 2
func leastMinutes(n int) int {
	return bits.Len(uint(n)-1)+1
}

# 3
func leastMinutes(n int) int {
	res := int(math.Ceil(math.Log(float64(n)) / math.Log(2)))
	return res + 1
}
```

## LCS02.完成一半题目(1)

- 题目

```
有 N 位扣友参加了微软与力扣举办了「以扣会友」线下活动。
主办方提供了 2*N 道题目，整型数组 questions 中每个数字对应了每道题目所涉及的知识点类型。
若每位扣友选择不同的一题，请返回被选的 N 道题目至少包含多少种知识点类型。
示例 1：输入：questions = [2,1,6,2] 输出：1
解释：有 2 位扣友在 4 道题目中选择 2 题。
可选择完成知识点类型为 2 的题目时，此时仅一种知识点类型
因此至少包含 1 种知识点类型。
示例 2：输入：questions = [1,5,1,3,4,5,2,5,3,3,8,6] 输出：2 
解释：有 6 位扣友在 12 道题目中选择题目，需要选择 6 题。
选择完成知识点类型为 3、5 的题目，因此至少包含 2 种知识点类型。
提示：questions.length == 2*n
2 <= questions.length <= 10^5
1 <= questions[i] <= 1000
```

- 解题思路

| No. | 思路 | 时间复杂度      | 空间复杂度 |
|:----|:---|:-----------|:------|
| 01  | 遍历 | O(nlog(n)) | O(n)  |

```go
func halfQuestions(questions []int) int {
	res := 0
	n := len(questions)
	arr := make([]int, 1001)
	for i := 0; i < n; i++ {
		arr[questions[i]]++
	}
	sort.Ints(arr)
	count := n / 2
	for i := 1000; i >= 0; i-- {
		res++
		count = count - arr[i]
		if count <= 0 {
			break
		}
	}
	return res
}
```

## LCS03.主题空间(1)

- 题目

```
「以扣会友」线下活动所在场地由若干主题空间与走廊组成，场地的地图记作由一维字符串型数组 grid，
字符串中仅包含 "0"～"5" 这 6 个字符。
地图上每一个字符代表面积为 1 的区域，其中 "0" 表示走廊，其他字符表示主题空间。
相同且连续（连续指上、下、左、右四个方向连接）的字符组成同一个主题空间。
假如整个 grid 区域的外侧均为走廊。请问，不与走廊直接相邻的主题空间的最大面积是多少？如果不存在这样的空间请返回 0。
示例 1:输入：grid = ["110","231","221"] 输出：1
解释：4 个主题空间中，只有 1 个不与走廊相邻，面积为 1。
示例 2:输入：grid = ["11111100000","21243101111","21224101221","11111101111"] 输出：3
解释：8 个主题空间中，有 5 个不与走廊相邻，面积分别为 3、1、1、1、2，最大面积为 3。
提示：1 <= grid.length <= 500
1 <= grid[i].length <= 500
grid[i][j] 仅可能是 "0"～"5"
```

- 解题思路

| No. | 思路     | 时间复杂度  | 空间复杂度 |
|:----|:-------|:-------|:------|
| 01  | 深度优先搜索 | O(n^2) | O(n)  |

```go
var count int
var flag bool

func largestArea(grid []string) int {
	res := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if '1' <= grid[i][j] && grid[i][j] <= '5' {
				count, flag = 0, true
				dfs(grid, i, j, grid[i][j])
				if flag == true && count > res {
					res = count
				}
			}
		}
	}
	return res
}

func dfs(grid []string, i, j int, char byte) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' { // 不满足条件
		flag = false
		return
	}
	if grid[i][j] != char { // 不一致
		return
	}
	count++
	arr := []byte(grid[i])
	arr[j] = arr[j] + 5 // 变换为其它的数
	grid[i] = string(arr)
	dfs(grid, i+1, j, char)
	dfs(grid, i-1, j, char)
	dfs(grid, i, j+1, char)
	dfs(grid, i, j-1, char)
	return
}
```