# LCP

* [LCP](#lcp)
	* [LCP01.猜数字(2)](#lcp01猜数字2)
	* [LCP02.分式化简(2)](#lcp02分式化简2)
	* [LCP06.拿硬币(2)](#lcp06拿硬币2)
	* [LCP07.传递信息(5)](#lcp07传递信息5)
	* [LCP11.期望个数统计(2)](#lcp11期望个数统计2)

## LCP01.猜数字(2)

- 题目

```
小A 和 小B 在玩猜数字。小B 每次从 1, 2, 3 中随机选择一个，小A 每次也从 1, 2, 3 中选择一个猜。
他们一共进行三次这个游戏，请返回 小A 猜对了几次？
输入的guess数组为 小A 每次的猜测，answer数组为 小B 每次的选择。
guess和answer的长度都等于3。
示例 1：输入：guess = [1,2,3], answer = [1,2,3] 输出：3
解释：小A 每次都猜对了。
示例 2：输入：guess = [2,2,3], answer = [3,2,1] 输出：1
解释：小A 只猜对了第二次。
限制：
    guess的长度 = 3
    answer的长度 = 3
    guess的元素取值为 {1, 2, 3} 之一。
    answer的元素取值为 {1, 2, 3} 之一。
```

- 解题思路

| No.  | 思路        | 时间复杂度 | 空间复杂度 |
| ---- | ----------- | ---------- | ---------- |
| 01   | 遍历        | O(1)       | O(1)       |
| 02   | 位运算-异或 | O(1)       | O(1)       |

```go
func game(guess []int, answer []int) int {
	res := 0
	for i := 0; i < len(guess); i++ {
		if guess[i] == answer[i] {
			res++
		}
	}
	return res
}

#
func game(guess []int, answer []int) int {
	res := 0
	for i := 0; i < len(guess); i++ {
		if guess[i]^answer[i] == 0 {
			res++
		}
	}
	return res
}
```

## LCP02.分式化简(2)

- 题目

```
有一个同学在学习分式。他需要将一个连分数化成最简分数，你能帮助他吗？
连分数是形如上图的分式。在本题中，所有系数都是大于等于0的整数。
输入的cont代表连分数的系数（cont[0]代表上图的a0，以此类推）。
返回一个长度为2的数组[n, m]，使得连分数的值等于n / m，且n, m最大公约数为1。
示例 1：输入：cont = [3, 2, 0, 2] 输出：[13, 4]
解释：原连分数等价于3 + (1 / (2 + (1 / (0 + 1 / 2))))。注意[26, 8], [-13, -4]都不是正确答案。
示例 2：输入：cont = [0, 0, 3] 输出：[3, 1]
解释：如果答案是整数，令分母为1即可。
限制：
    cont[i] >= 0
    1 <= cont的长度 <= 10
    cont最后一个元素不等于0
    答案的n, m的取值都能被32位int整型存下（即不超过2 ^ 31 - 1）。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |
| 02   | 递归 | O(n)       | O(n)       |

```go
func fraction(cont []int) []int {
	n, m := 1, cont[len(cont)-1]
	for i := len(cont) - 2; i >= 0; i-- {
		n, m = m, cont[i]*m+n
	}
	return []int{m, n}
}

#
func fraction(cont []int) []int {
	if len(cont) == 1 {
		return []int{cont[0], 1}
	}
	n := fraction(cont[1:])
	m := cont[0]
	return []int{m*n[0] + n[1], n[0]}
}
```

## LCP06.拿硬币(2)

- 题目

```
桌上有 n 堆力扣币，每堆的数量保存在数组 coins 中。
我们每次可以选择任意一堆，拿走其中的一枚或者两枚，求拿完所有力扣币的最少次数。
示例 1：输入：[4,2,1]输出：4
解释：第一堆力扣币最少需要拿 2 次，第二堆最少需要拿 1 次，第三堆最少需要拿 1 次，总共 4 次即可拿完。
示例 2： 输入：[2,3,10]输出：8
限制：
    1 <= n <= 4
    1 <= coins[i] <= 10
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 遍历     | O(n)       | O(1)       |
| 02   | 内置函数 | O(n)       | O(1)       |

```go
func minCount(coins []int) int {
	res := 0
	for i := 0; i < len(coins); i++ {
		res = res + coins[i]/2
		if coins[i]%2 == 1 {
			res = res + 1
		}
	}
	return res
}

#
func minCount(coins []int) int {
	res := 0
	for i := 0; i < len(coins); i++ {
		res = res + int(math.Ceil(float64(coins[i])/2))
	}
	return res
}
```

## LCP07.传递信息(5)

- 题目

```
小朋友 A 在和 ta 的小伙伴们玩传信息游戏，游戏规则如下：
    有 n 名玩家，所有玩家编号分别为 0 ～ n-1，其中小朋友 A 的编号为 0
    每个玩家都有固定的若干个可传信息的其他玩家（也可能没有）。
    传信息的关系是单向的（比如 A 可以向 B 传信息，但 B 不能向 A 传信息）。
    每轮信息必须需要传递给另一个人，且信息可重复经过同一个人
给定总玩家数 n，以及按 [玩家编号,对应可传递玩家编号] 关系组成的二维数组 relation。
返回信息从小 A (编号 0 ) 经过 k 轮传递到编号为 n-1 的小伙伴处的方案数；若不能到达，返回 0。

示例 1：
    输入：n = 5, relation = [[0,2],[2,1],[3,4],[2,3],[1,4],[2,0],[0,4]], k = 3
    输出：3
    解释：信息从小 A 编号 0 处开始，经 3 轮传递，到达编号 4。
    共有 3 种方案，分别是 0->2->0->4， 0->2->1->4， 0->2->3->4。
示例 2：
    输入：n = 3, relation = [[0,2],[2,1]], k = 2
    输出：0
    解释：信息不能从小 A 处经过 2 轮传递到编号 2
限制：
    2 <= n <= 10
    1 <= k <= 5
    1 <= relation.length <= 90, 且 relation[i].length == 2
    0 <= relation[i][0],relation[i][1] < n 且 relation[i][0] != relation[i][1]
```

- 解题思路

| No.  | 思路         | 时间复杂度 | 空间复杂度 |
| ---- | ------------ | ---------- | ---------- |
| 01   | 深度优先搜索 | O(n^k)     | O(n^2)     |
| 02   | 深度优先搜索 | O(n^k)     | O(n)       |
| 03   | 广度优先搜索 | O(n^k)     | O(n^k)     |
| 04   | 动态规划     | O(n^2)     | O(n)       |
| 05   | 动态规划     | O(n^2)     | O(n^2)     |

```go
var ways [][]bool

func numWays(n int, relation [][]int, k int) int {
	ways = make([][]bool, n)
	for i := range ways {
		ways[i] = make([]bool, n)
	}
	sum := 0
	for i := 0; i < len(relation); i++ {
		ways[relation[i][0]][relation[i][1]] = true
	}
	for i := 0; i < n; i++ {
		if ways[0][i] == true {
			sum = sum + dfs(i, n, k-1)
		}
	}
	return sum
}

func dfs(i, n, k int) int {
	sum := 0
	if k < 0 || i < 0 || i >= n {
		return 0
	}
	if k == 0 && i == n-1 {
		return 1
	} else {
		for j := 0; j < n; j++ {
			if ways[i][j] == true {
				sum += dfs(j, n, k-1)
			}
		}
	}
	return sum
}

#
var res, K, N int

func numWays(n int, relation [][]int, k int) int {
	res = 0
	K = k
	N = n
	dfs(0, relation, 0)
	return res
}

func dfs(n int, relation [][]int, k int) {
	if n == N-1 && k == K {
		res++
	}
	if k > K {
		return
	}
	for i := 0; i < len(relation); i++ {
		if relation[i][0] == n {
			dfs(relation[i][1], relation, k+1)
		}
	}
}

#
func numWays(n int, relation [][]int, k int) int {
	m := make(map[int][]int)
	for i := 0; i < len(relation); i++ {
		m[relation[i][0]] = append(m[relation[i][0]], relation[i][1])
	}
	start := []int{0}
	for i := 0; i < k; i++ {
		arr := make([]int, 0)
		for j := 0; j < len(start); j++ {
			for k := 0; k < len(m[start[j]]); k++ {
				arr = append(arr, m[start[j]][k])
			}
		}
		start = arr
	}
	res := 0
	for i := 0; i < len(start); i++ {
		if start[i] == n-1 {
			res++
		}
	}
	return res
}

# 4
func numWays(n int, relation [][]int, k int) int {
	dp := make([]int, n)
	dp[0] = 1
	for i := 0; i < k; i++ {
		temp := make([]int, n)
		for _, v := range relation {
			temp[v[1]] = temp[v[1]] + dp[v[0]]
		}
		dp = temp
	}
	return dp[n-1]
}

# 5
func numWays(n int, relation [][]int, k int) int {
	dp := make([][]int, k+1)
	for i := 0; i < k+1; i++ {
		dp[i] = make([]int, n)
	}
	dp[0][0] = 1
	for i := 0; i < k; i++ {
		for _, v := range relation {
			dp[i+1][v[1]] = dp[i+1][v[1]] + dp[i][v[0]]
		}
	}
	return dp[k][n-1]
}
```

## LCP11.期望个数统计(2)

- 题目

```
某互联网公司一年一度的春招开始了，一共有 n 名面试者入选。
每名面试者都会提交一份简历，公司会根据提供的简历资料产生一个预估的能力值，数值越大代表越有可能通过面试。
小 A 和小 B 负责审核面试者，他们均有所有面试者的简历，并且将各自根据面试者能力值从大到小的顺序浏览。
由于简历事先被打乱过，能力值相同的简历的出现顺序是从它们的全排列中等可能地取一个。
现在给定 n 名面试者的能力值 scores，
设 X 代表小 A 和小 B 的浏览顺序中出现在同一位置的简历数，求 X 的期望。
提示：离散的非负随机变量的期望计算公式为 1。在本题中，由于 X 的取值为 0 到 n 之间，期望计算公式可以是 2。
示例 1： 输入：scores = [1,2,3] 输出：3
    解释：由于面试者能力值互不相同，小 A 和小 B 的浏览顺序一定是相同的。X的期望是 3 。
示例 2：输入：scores = [1,1]  输出：1 解释：设两位面试者的编号为 0, 1。
    由于他们的能力值都是 1，小 A 和小 B 的浏览顺序都为从全排列 [[0,1],[1,0]] 中等可能地取一个。
    如果小 A 和小 B 的浏览顺序都是 [0,1] 或者 [1,0] ，那么出现在同一位置的简历数为 2 ，否则是 0 。
    所以 X 的期望是 (2+0+2+0) * 1/4 = 1
示例 3：输入：scores = [1,1,2] 输出：2
限制：
    1 <= scores.length <= 10^5
    0 <= scores[i] <= 10^6
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n)       | O(n)       |
| 02   | 排序遍历 | O(nlog(n)) | O(1) |

```go
func expectNumber(scores []int) int {
	m := make(map[int]bool)
	for i := 0; i < len(scores); i++{
		m[scores[i]] = true
	}
	return len(m)
}

#
func expectNumber(scores []int) int {
	sort.Ints(scores)
	count := 0
	for i := 1; i < len(scores); i++ {
		if scores[i] == scores[i-1] {
			count++
		}
	}
	return len(scores) - count
}
```

