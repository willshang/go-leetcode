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

## LCP03.机器人大冒险(1)

- 题目

```
力扣团队买了一个可编程机器人，机器人初始位置在原点(0, 0)。
小伙伴事先给机器人输入一串指令command，机器人就会无限循环这条指令的步骤进行移动。指令有两种：
    U: 向y轴正方向移动一格
    R: 向x轴正方向移动一格。
不幸的是，在 xy 平面上还有一些障碍物，他们的坐标用obstacles表示。机器人一旦碰到障碍物就会被损毁。
给定终点坐标(x, y)，返回机器人能否完好地到达终点。如果能，返回true；否则返回false。
示例 1：输入：command = "URR", obstacles = [], x = 3, y = 2 输出：true
解释：U(0, 1) -> R(1, 1) -> R(2, 1) -> U(2, 2) -> R(3, 2)。
示例 2：输入：command = "URR", obstacles = [[2, 2]], x = 3, y = 2 输出：false
解释：机器人在到达终点前会碰到(2, 2)的障碍物。
示例 3：输入：command = "URR", obstacles = [[4, 2]], x = 3, y = 2 输出：true
解释：到达终点后，再碰到障碍物也不影响返回结果。
限制：
    2 <= command的长度 <= 1000
    command由U，R构成，且至少有一个U，至少有一个R
    0 <= x <= 1e9, 0 <= y <= 1e9
    0 <= obstacles的长度 <= 1000
    obstacles[i]不为原点或者终点
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 数学计算 | O(n)       | O(1)       |

```go
func robot(command string, obstacles [][]int, x int, y int) bool {
	if judge(command, x, y) == false {
		return false
	}
	for _, node := range obstacles {
		if x >= node[0] && y >= node[1] && judge(command, node[0], node[1]) {
			return false
		}
	}
	return true
}

func judge(command string, x, y int) bool {
	u := strings.Count(command, "U")
	r := strings.Count(command, "R")
	times := (x + y) / len(command)
	last := command[:(x+y)%len(command)]
	uNum := u*times + strings.Count(last, "U")
	rNum := r*times + strings.Count(last, "R")
	if uNum == y && rNum == x {
		return true
	}
	return false
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

# 3
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

## LCP08.剧情触发时间(2)

- 题目

```
在战略游戏中，玩家往往需要发展自己的势力来触发各种新的剧情。一个势力的主要属性有三种，
分别是文明等级（C），资源储备（R）以及人口数量（H）。在游戏开始时（第 0 天），三种属性的值均为 0。
随着游戏进程的进行，每一天玩家的三种属性都会对应增加，我们用一个二维数组 increase 来表示每天的增加情况。
这个二维数组的每个元素是一个长度为 3 的一维数组，
例如 [[1,2,1],[3,4,2]] 表示第一天三种属性分别增加 1,2,1 而第二天分别增加 3,4,2。
所有剧情的触发条件也用一个二维数组 requirements 表示。
这个二维数组的每个元素是一个长度为 3 的一维数组，对于某个剧情的触发条件 c[i], r[i], h[i]，
如果当前 C >= c[i] 且 R >= r[i] 且 H >= h[i] ，则剧情会被触发。
根据所给信息，请计算每个剧情的触发时间，并以一个数组返回。
如果某个剧情不会被触发，则该剧情对应的触发时间为 -1 。
示例 1：
    输入： increase = [[2,8,4],[2,5,0],[10,9,8]] 
    requirements = [[2,11,3],[15,10,7],[9,17,12],[8,1,14]]
    输出: [2,-1,3,-1]
    解释：
    初始时，C = 0，R = 0，H = 0
    第 1 天，C = 2，R = 8，H = 4
    第 2 天，C = 4，R = 13，H = 4，此时触发剧情 0
    第 3 天，C = 14，R = 22，H = 12，此时触发剧情 2
    剧情 1 和 3 无法触发。
示例 2：
    输入： increase = [[0,4,5],[4,8,8],[8,6,1],[10,10,0]] 
    requirements = [[12,11,16],[20,2,6],[9,2,6],[10,18,3],[8,14,9]]
    输出: [-1,4,3,3,3]
示例 3：输入： increase = [[1,1,1]] requirements = [[0,0,0]] 输出: [0]
限制：
    1 <= increase.length <= 10000
    1 <= requirements.length <= 100000
    0 <= increase[i] <= 10
    0 <= requirements[i] <= 100000
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 二分查找 | O(nlog(n)) | O(n)       |
| 02   | 内置函数 | O(nlog(n)) | O(n)       |

```go
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	for i := 1; i < len(increase); i++ {
		increase[i][0] = increase[i][0] + increase[i-1][0]
		increase[i][1] = increase[i][1] + increase[i-1][1]
		increase[i][2] = increase[i][2] + increase[i-1][2]
	}
	res := make([]int, len(requirements))
	for i := 0; i < len(requirements); i++ {
		C, R, H := requirements[i][0], requirements[i][1], requirements[i][2]
		if C == 0 && R == 0 && H == 0 {
			res[i] = 0
			continue
		}
		if C > increase[len(increase)-1][0] ||
			R > increase[len(increase)-1][1] ||
			H > increase[len(increase)-1][2] {
			res[i] = -1
			continue
		}
		left, right := 0, len(increase)-1
		index := -1
		for left <= right {
			mid := left + (right-left)/2
			if increase[mid][0] >= C && increase[mid][1] >= R && increase[mid][2] >= H {
				index = mid + 1
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
		res[i] = index
	}
	return res
}

#
func getTriggerTime(increase [][]int, requirements [][]int) []int {
	for i := 1; i < len(increase); i++ {
		increase[i][0] = increase[i][0] + increase[i-1][0]
		increase[i][1] = increase[i][1] + increase[i-1][1]
		increase[i][2] = increase[i][2] + increase[i-1][2]
	}

	res := make([]int, len(requirements))
	for i := 0; i < len(requirements); i++ {
		C, R, H := requirements[i][0], requirements[i][1], requirements[i][2]
		if C == 0 && R == 0 && H == 0 {
			res[i] = 0
			continue
		}
		if C > increase[len(increase)-1][0] ||
			R > increase[len(increase)-1][1] ||
			H > increase[len(increase)-1][2] {
			res[i] = -1
			continue
		}
		index := sort.Search(len(increase), func(j int) bool {
			return increase[j][0] >= requirements[i][0] &&
				increase[j][1] >= requirements[i][1] &&
				increase[j][2] >= requirements[i][2]
		})
		if index == len(increase) {
			index = -2
		}
		res[i] = index + 1
	}
	return res
}
```

## LCP09.最小跳跃次数

### 题目

```
为了给刷题的同学一些奖励，力扣团队引入了一个弹簧游戏机。
游戏机由 N 个特殊弹簧排成一排，编号为 0 到 N-1。初始有一个小球在编号 0 的弹簧处。
若小球在编号为 i 的弹簧处，通过按动弹簧，可以选择把小球向右弹射 jump[i] 的距离，或者向左弹射到任意左侧弹簧的位置。
也就是说，在编号为 i 弹簧处按动弹簧，小球可以弹向 0 到 i-1 中任意弹簧或者 i+jump[i] 的弹簧（若 i+jump[i]>=N ，
则表示小球弹出了机器）。小球位于编号 0 处的弹簧时不能再向左弹。
为了获得奖励，你需要将小球弹出机器。
请求出最少需要按动多少次弹簧，可以将小球从编号 0 弹簧弹出整个机器，即向右越过编号 N-1 的弹簧。
示例 1：输入：jump = [2, 5, 1, 1, 1, 1]输出：3
解释：小 Z 最少需要按动 3 次弹簧，小球依次到达的顺序为 0 -> 2 -> 1 -> 6，最终小球弹出了机器。
限制：1 <= jump.length <= 10^6
1 <= jump[i] <= 10000
```

### 解题思路

```go

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

## LCP12.小张刷题计划(2)

- 题目

```
为了提高自己的代码能力，小张制定了 LeetCode 刷题计划，他选中了 LeetCode 题库中的 n 道题，
编号从 0 到 n-1，并计划在 m 天内按照题目编号顺序刷完所有的题目（注意，小张不能用多天完成同一题）。
在小张刷题计划中，小张需要用 time[i] 的时间完成编号 i 的题目。此外，小张还可以使用场外求助功能，
通过询问他的好朋友小杨题目的解法，可以省去该题的做题时间。为了防止“小张刷题计划”变成“小杨刷题计划”，
小张每天最多使用一次求助。
我们定义 m 天中做题时间最多的一天耗时为 T（小杨完成的题目不计入做题总时间）。
请你帮小张求出最小的 T是多少。
示例 1：
    输入：time = [1,2,3,3], m = 2
    输出：3
    解释：第一天小张完成前三题，其中第三题找小杨帮忙；第二天完成第四题，并且找小杨帮忙。
    这样做题时间最多的一天花费了 3 的时间，并且这个值是最小的。
示例 2：输入：time = [999,999,999], m = 4 输出：0
    解释：在前三天中，小张每天求助小杨一次，这样他可以在三天内完成所有的题目并不花任何时间。
限制：
    1 <= time.length <= 10^5
    1 <= time[i] <= 10000
    1 <= m <= 1000
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 二分查找 | O(nlog(n)) | O(1)       |
| 02   | 二分查找 | O(nlog(n)) | O(1)       |

```go
func minTime(time []int, m int) int {
	left, right, mid := 0, 0, 0
	for i := 0; i < len(time); i++ {
		right = right + time[i]
	}
	// 二分查找一个数mid，使time数组能分割成m个和不小于mid的子数组
	for left <= right {
		mid = left + (right-left)/2
		if check(time, mid, m) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func check(arr []int, mid, m int) bool {
	maxValue := 0
	sum := 0
	count := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if sum-maxValue > mid {
			count++
			if count >= m {
				return false
			}
			sum = arr[i]
			maxValue = arr[i]
		}
	}
	return true
}

#
func minTime(time []int, m int) int {
	left, right, mid := 0, 0, 0
	for i := 0; i < len(time); i++ {
		right = right + time[i]
	}
	// 二分查找一个数mid，使time数组能分割成m个和不小于mid的子数组
	res := math.MaxInt32
	for left <= right {
		mid = left + (right-left)/2
		if check(time, mid) <= m {
			if mid < res {
				res = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}

func check(arr []int, mid int) int {
	res := 1
	maxValue := 0
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
		if arr[i] > maxValue {
			maxValue = arr[i]
		}
		if sum-maxValue > mid {
			sum = arr[i]
			maxValue = arr[i]
			res++
		}
	}
	return res
}
```

## LCP17.速算机器人(2)

- 题目

```
小扣在秋日市集发现了一款速算机器人。店家对机器人说出两个数字（记作 x 和 y），请小扣说出计算指令：
    "A" 运算：使 x = 2 * x + y；
    "B" 运算：使 y = 2 * y + x。
在本次游戏中，店家说出的数字为 x = 1 和 y = 0，小扣说出的计算指令记作仅由大写字母 A、B 组成的字符串 s，
字符串中字符的顺序表示计算顺序，请返回最终 x 与 y 的和为多少。
示例 1： 输入：s = "AB" 输出：4
解释：经过一次 A 运算后，x = 2, y = 0。
    再经过一次 B 运算，x = 2, y = 2。
    最终 x 与 y 之和为 4。
提示：0 <= s.length <= 10
    s 由 'A' 和 'B' 组成
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 遍历 | O(n)       | O(1)       |
| 02   | 数学 | O(1)       | O(1)       |

```go
func calculate(s string) int {
	x, y := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == 'A' {
			x = 2*x + y
		} else if s[i] == 'B' {
			y = 2*y + x
		}
	}
	return x + y
}

# 2
func calculate(s string) int {
	return 1 << len(s)
}
```

## LCP18.早餐组合(3)

- 题目

```
小扣在秋日市集选择了一家早餐摊位，一维整型数组 staple 中记录了每种主食的价格，
一维整型数组 drinks 中记录了每种饮料的价格。小扣的计划选择一份主食和一款饮料，且花费不超过 x 元。
请返回小扣共有多少种购买方案。
注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1
示例 1：输入：staple = [10,20,5], drinks = [5,5,2], x = 15 输出：6
    解释：小扣有 6 种购买方案，所选主食与所选饮料在数组中对应的下标分别是：
    第 1 种方案：staple[0] + drinks[0] = 10 + 5 = 15；
    第 2 种方案：staple[0] + drinks[1] = 10 + 5 = 15；
    第 3 种方案：staple[0] + drinks[2] = 10 + 2 = 12；
    第 4 种方案：staple[2] + drinks[0] = 5 + 5 = 10；
    第 5 种方案：staple[2] + drinks[1] = 5 + 5 = 10；
    第 6 种方案：staple[2] + drinks[2] = 5 + 2 = 7。
示例 2： 输入：staple = [2,1,1], drinks = [8,9,5,1], x = 9 输出：8
    解释：小扣有 8 种购买方案，所选主食与所选饮料在数组中对应的下标分别是：
    第 1 种方案：staple[0] + drinks[2] = 2 + 5 = 7；
    第 2 种方案：staple[0] + drinks[3] = 2 + 1 = 3；
    第 3 种方案：staple[1] + drinks[0] = 1 + 8 = 9；
    第 4 种方案：staple[1] + drinks[2] = 1 + 5 = 6；
    第 5 种方案：staple[1] + drinks[3] = 1 + 1 = 2；
    第 6 种方案：staple[2] + drinks[0] = 1 + 8 = 9；
    第 7 种方案：staple[2] + drinks[2] = 1 + 5 = 6；
    第 8 种方案：staple[2] + drinks[3] = 1 + 1 = 2；
提示：1 <= staple.length <= 10^5
    1 <= drinks.length <= 10^5
    1 <= staple[i],drinks[i] <= 10^5
    1 <= x <= 2*10^5
```

- 解题思路

| No.  | 思路            | 时间复杂度 | 空间复杂度 |
| ---- | --------------- | ---------- | ---------- |
| 01   | 排序双指针      | O(nlog(n)) | O(1)       |
| 02   | 数组辅助+前缀和 | O(n)       | O(n)       |
| 03   | 排序+二分查找   | O(nlog(n)) | O(1)       |

```go
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	res := 0
	j := len(drinks) - 1
	for i := 0; i < len(staple); i++ {
		for j >= 0 && staple[i]+drinks[j] > x {
			j--
		}
		res = (res + j + 1) % 1000000007
	}
	return res
}

# 2
func breakfastNumber(staple []int, drinks []int, x int) int {
	res := 0
	arr := make([]int, x+1)
	for i := 0; i < len(staple); i++ {
		if staple[i] < x {
			arr[staple[i]]++
		}
	}
	for i := 1; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i]
	}
	for i := 0; i < len(drinks); i++ {
		target := x - drinks[i]
		if target <= 0 {
			continue
		}
		res = (res + arr[target]) % 1000000007
	}
	return res
}

# 3
func breakfastNumber(staple []int, drinks []int, x int) int {
	sort.Ints(staple)
	sort.Ints(drinks)
	res := 0
	for i := 0; i < len(staple); i++ {
		target := x - staple[i]
		if target <= 0 {
			break
		}
		j := binarySearch(drinks, target)
		res = (res + j) % 1000000007
	}
	return res
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := left + (right-left)/2
		if arr[mid] > target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```

## LCP19.秋叶收藏集(2)

- 题目

```
小扣出去秋游，途中收集了一些红叶和黄叶，他利用这些叶子初步整理了一份秋叶收藏集 leaves， 
字符串 leaves 仅包含小写字符 r 和 y， 其中字符 r 表示一片红叶，字符 y 表示一片黄叶。
出于美观整齐的考虑，小扣想要将收藏集中树叶的排列调整成「红、黄、红」三部分。
每部分树叶数量可以不相等，但均需大于等于 1。每次调整操作，
小扣可以将一片红叶替换成黄叶或者将一片黄叶替换成红叶。
请问小扣最少需要多少次调整操作才能将秋叶收藏集调整完毕。
示例 1：输入：leaves = "rrryyyrryyyrr" 输出：2
解释：调整两次，将中间的两片红叶替换成黄叶，得到 "rrryyyyyyyyrr"
示例 2：输入：leaves = "ryr" 输出：0
解释：已符合要求，不需要额外操作
提示：3 <= leaves.length <= 10^5
    leaves 中只包含字符 'r' 和字符 'y'
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 动态规划-二维 | O(n)       | O(n)       |
| 02   | 动态规划-二维 | O(n)       | O(n)       |

```go
func minimumOperations(leaves string) int {
	n := len(leaves)
	// 长度i+1
	// dp[i][0] 全部变成r的步数
	// dp[i][1] 变成r...ry...y的步数
	// dp[i][2] 变成r...ry...yr...r的步数
	dp := make([][3]int, n)
	if leaves[0] == 'y' {
		dp[0][0] = 1 // 1个y变为r需要1步
	}
	for i := 1; i < n; i++ {
		if leaves[i] == 'r' {
			dp[i][0] = dp[i-1][0]     // 不需要改变，同前一个
			dp[i][1] = dp[i-1][0] + 1 // 全r + 当前r，需要改变一个y，步数+1
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
				dp[i][2] = dp[i-1][1]
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2])
			}
		} else {
			dp[i][0] = dp[i-1][0] + 1 // 需要改变，步数+1
			dp[i][1] = dp[i-1][0]     // 前一个全r + 当前y,不需要改变
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1]) // 同前一个不变
				dp[i][2] = dp[i-1][1] + 1            // 调整+1
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2]+1) // 在前一个基础上调整
			}
		}
	}
	return dp[n-1][2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 2
func minimumOperations(leaves string) int {
	n := len(leaves)
	// 长度i+1
	// dp[i][0] 全部变成r的步数
	// dp[i][1] 变成r...ry...y的步数
	// dp[i][2] 变成r...ry...yr...r的步数
	dp := make([][3]int, n)
	maxValue := math.MaxInt32 / 10
	if leaves[0] == 'y' {
		dp[0][0] = 1 // 1个y变为r需要1步
		dp[0][1] = maxValue
		dp[0][2] = maxValue
	}
	for i := 1; i < n; i++ {
		dp[i][2] = maxValue
		dp[i][1] = maxValue
		if leaves[i] == 'r' {
			dp[i][0] = dp[i-1][0]                      // 不需要改变，同前一个
			dp[i][1] = min(dp[i-1][0]+1, dp[i-1][1]+1) // 前一个r+1，前一个y+1
			if i >= 2 {
				dp[i][2] = min(dp[i-1][1], dp[i-1][2]) // 前一个y不变，前一个r不变
			}
		} else {
			dp[i][0] = dp[i-1][0] + 1              // 需要改变，步数+1
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) // 前一个r不变，前一个y不变
			if i >= 2 {
				dp[i][2] = min(dp[i-1][1]+1, dp[i-1][2]+1) // 前一个y+1, 前一个r+1
			}
		}
	}
	return dp[n-1][2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

## LCP20.快速公交(1)

- 题目

```
小扣打算去秋日市集，由于游客较多，小扣的移动速度受到了人流影响：
小扣从 x 号站点移动至 x + 1 号站点需要花费的时间为 inc；
小扣从 x 号站点移动至 x - 1 号站点需要花费的时间为 dec。
现有 m 辆公交车，编号为 0 到 m-1。小扣也可以通过搭乘编号为 i 的公交车，从 x 号站点移动至 jump[i]*x 号站点，
耗时仅为 cost[i]。小扣可以搭乘任意编号的公交车且搭乘公交次数不限。
假定小扣起始站点记作 0，秋日市集站点记作 target，请返回小扣抵达秋日市集最少需要花费多少时间。
由于数字较大，最终答案需要对 1000000007 (1e9 + 7) 取模。
注意：小扣可在移动过程中到达编号大于 target 的站点。
示例 1：输入：target = 31, inc = 5, dec = 3, jump = [6], cost = [10] 输出：33
解释：小扣步行至 1 号站点，花费时间为 5；
小扣从 1 号站台搭乘 0 号公交至 6 * 1 = 6 站台，花费时间为 10；
小扣从 6 号站台步行至 5 号站台，花费时间为 3；
小扣从 5 号站台搭乘 0 号公交至 6 * 5 = 30 站台，花费时间为 10；
小扣从 30 号站台步行至 31 号站台，花费时间为 5；
最终小扣花费总时间为 33。
示例 2：输入：target = 612, inc = 4, dec = 5, jump = [3,6,8,11,5,10,4], cost = [4,7,6,3,7,6,4] 输出：26
解释：小扣步行至 1 号站点，花费时间为 4；
小扣从 1 号站台搭乘 0 号公交至 3 * 1 = 3 站台，花费时间为 4；
小扣从 3 号站台搭乘 3 号公交至 11 * 3 = 33 站台，花费时间为 3；
小扣从 33 号站台步行至 34 站台，花费时间为 4；
小扣从 34 号站台搭乘 0 号公交至 3 * 34 = 102 站台，花费时间为 4；
小扣从 102 号站台搭乘 1 号公交至 6 * 102 = 612 站台，花费时间为 7；
最终小扣花费总时间为 26。
提示：1 <= target <= 10^9
1 <= jump.length, cost.length <= 10
2 <= jump[i] <= 10^6
1 <= inc, dec, cost[i] <= 10^6
```

- 解题思路

| No.  | 思路            | 时间复杂度   | 空间复杂度 |
| ---- | --------------- | ------------ | ---------- |
| 01   | 堆+广度优先搜索 | O(n^2log(n)) | O(n)       |
| 02   | 深度优先搜索    | O(n^2)       | O(n)       |

```go
var mod = 1000000007

var res int

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	res = target * inc // 最坏的情况：全+1
	n := len(jump)
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	visited := make(map[int]bool)
	heap.Push(&intHeap, []int{0, target}) // 时间+当前位置：从后往前走
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([]int)
		t, cur := node[0], node[1]
		if t >= res { // 跳过
			continue
		}
		res = min(res, t+cur*inc) // 用+1补
		for i := 0; i < n; i++ {
			diff, next := cur%jump[i], cur/jump[i]
			if diff == 0 { // 直接坐公交
				if visited[next] == false {
					heap.Push(&intHeap, []int{t + cost[i], next})
				}
			} else {
				if visited[next] == false { // 向左走坐公交
					heap.Push(&intHeap, []int{t + cost[i] + diff*inc, next})
				}
				if visited[next+1] == false { // 向右走坐公交
					heap.Push(&intHeap, []int{t + cost[i] + (jump[i]-diff)*dec, next + 1})
				}
			}
		}
	}
	return res % mod
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

# 2
var mod = 1000000007

var visited map[int]int

func busRapidTransit(target int, inc int, dec int, jump []int, cost []int) int {
	visited = make(map[int]int)
	return dfs(inc, dec, jump, cost, target) % mod
}

func dfs(inc int, dec int, jump []int, cost []int, cur int) int {
	if cur == 0 {
		return 0
	}
	if cur == 1 {
		return inc
	}
	if visited[cur] > 0 {
		return visited[cur]
	}
	res := cur * inc // 最坏的情况：全+1
	for i := 0; i < len(jump); i++ {
		diff, next := cur%jump[i], cur/jump[i]
		if diff == 0 { // 直接坐公交
			res = min(res, dfs(inc, dec, jump, cost, next)+cost[i])
		} else {
			// 向左走坐公交
			res = min(res, dfs(inc, dec, jump, cost, next)+cost[i]+diff*inc)
			// 向右走坐公交
			res = min(res, dfs(inc, dec, jump, cost, next+1)+cost[i]+(jump[i]-diff)*dec)
		}
	}
	visited[cur] = res
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

## LCP22.黑白方格画(1)

- 题目

```
小扣注意到秋日市集上有一个创作黑白方格画的摊位。摊主给每个顾客提供一个固定在墙上的白色画板，画板不能转动。
画板上有 n * n 的网格。绘画规则为，小扣可以选择任意多行以及任意多列的格子涂成黑色，
所选行数、列数均可为 0。
小扣希望最终的成品上需要有 k 个黑色格子，请返回小扣共有多少种涂色方案。
注意：两个方案中任意一个相同位置的格子颜色不同，就视为不同的方案。
示例 1：输入：n = 2, k = 2 输出：4
解释：一共有四种不同的方案：
第一种方案：涂第一列；
第二种方案：涂第二列；
第三种方案：涂第一行；
第四种方案：涂第二行。
示例 2：输入：n = 2, k = 1 输出：0
解释：不可行，因为第一次涂色至少会涂两个黑格。
示例 3：输入：n = 2, k = 4 输出：1
解释：共有 2*2=4 个格子，仅有一种涂色方案。
限制：1 <= n <= 6
0 <= k <= n * n
```

- 解题思路

| No.  | 思路        | 时间复杂度 | 空间复杂度 |
| ---- | ----------- | ---------- | ---------- |
| 01   | 暴力法+组合 | O(1)       | O(1)       |

```go
func paintingPlan(n int, k int) int {
	if k == n*n || k == 0 { // 全部涂满或者不涂只有1种方案
		return 1
	}
	if k < n { // 最少大于等于n
		return 0
	}
	res := 0
	// 暴力枚举行和列
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := i * n
			b := j * n
			if a+b-i*j == k {
				res = res + C(n, i)*C(n, j) // 求组合数
			}
		}
	}
	return res
}

func C(n, m int) int {
	a := 1
	for i := 1; i <= m; i++ {
		a = a * (n - i + 1)
	}

	b := 1
	for i := 1; i <= m; i++ {
		b = b * i
	}
	return a / b
}
```

## LCP23.魔术排列(1)

- 题目

```
秋日市集上，魔术师邀请小扣与他互动。魔术师的道具为分别写有数字 1~N 的 N 张卡牌，然后请小扣思考一个 N 张卡牌的排列 target。
魔术师的目标是找到一个数字 k（k >= 1），使得初始排列顺序为 1~N 的卡牌经过特殊的洗牌方式最终变成小扣所想的排列 target，
特殊的洗牌方式为：
第一步，魔术师将当前位于 偶数位置 的卡牌（下标自 1 开始），保持 当前排列顺序 放在位于 奇数位置 的卡牌之前。
例如：将当前排列 [1,2,3,4,5] 位于偶数位置的 [2,4] 置于奇数位置的 [1,3,5] 前，排列变为 [2,4,1,3,5]；
第二步，若当前卡牌数量小于等于 k，则魔术师按排列顺序取走全部卡牌；
若当前卡牌数量大于 k，则取走前 k 张卡牌，剩余卡牌继续重复这两个步骤，直至所有卡牌全部被取走；
卡牌按照魔术师取走顺序构成的新排列为「魔术取数排列」，请返回是否存在这个数字 k 使得「魔术取数排列」恰好就是 target，
从而让小扣感到大吃一惊。
示例 1：输入：target = [2,4,3,1,5] 输出：true
解释：排列 target 长度为 5，初始排列为：1,2,3,4,5。我们选择 k = 2：
第一次：将当前排列 [1,2,3,4,5] 位于偶数位置的 [2,4] 置于奇数位置的 [1,3,5] 前，排列变为 [2,4,1,3,5]。
取走前 2 张卡牌 2,4，剩余 [1,3,5]；
第二次：将当前排列 [1,3,5] 位于偶数位置的 [3] 置于奇数位置的 [1,5] 前，排列变为 [3,1,5]。
取走前 2 张 3,1，剩余 [5]；
第三次：当前排列为 [5]，全部取出。
最后，数字按照取出顺序构成的「魔术取数排列」2,4,3,1,5 恰好为 target。
示例 2：输入：target = [5,4,3,2,1] 出：false
解释：无法找到一个数字 k 可以使「魔术取数排列」恰好为 target。
提示：1 <= target.length = N <= 5000
题目保证 target 是 1~N 的一个排列。
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 模拟 | O(n^2)     | O(n)       |

```go
func isMagic(target []int) bool {
	n := len(target)
	arr := make([]int, n) // 构建初始数组
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	arr = Shuffle(arr)
	k := 0 // 找到k：可以证明只能正好匹配到k个；其中1~N不重复
	for ; k < n && arr[k] == target[k]; k++ {
	}
	if k == 0 { // 不满足
		return false
	}
	for { // 按规则模拟
		if k >= len(arr) {
			return judge(arr, target, len(arr))
		}
		if judge(arr, target, k) == false {
			return false
		}
		arr = arr[k:]
		arr = Shuffle(arr)
		target = target[k:]
	}
	return false
}

func Shuffle(arr []int) []int {
	temp := make([]int, 0)
	for i := 1; i < len(arr); i = i + 2 {
		temp = append(temp, arr[i])
	}
	for i := 0; i < len(arr); i = i + 2 {
		temp = append(temp, arr[i])
	}
	return temp
}

func judge(a, b []int, k int) bool {
	for i := 0; i < k; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

## LCP28.采购方案(3)

- 题目

```
小力将 N 个零件的报价存于数组 nums。小力预算为 target，假定小力仅购买两个零件，
要求购买零件的花费不超过预算，请问他有多少种采购方案。
注意：答案需要以 1e9 + 7 (1000000007) 为底取模，如：计算初始结果为：1000000008，请返回 1
示例 1：输入：nums = [2,5,3,5], target = 6 输出：1
解释：预算内仅能购买 nums[0] 与 nums[2]。
示例 2：输入：nums = [2,2,1,9], target = 10 输出：4
解释：符合预算的采购方案如下：
nums[0] + nums[1] = 4
nums[0] + nums[2] = 3
nums[1] + nums[2] = 3
nums[2] + nums[3] = 10
提示：2 <= nums.length <= 10^5
1 <= nums[i], target <= 10^5
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 排序+双指针   | O(nlog(n)) | O(1)       |
| 02   | 排序+双指针   | O(nlog(n)) | O(1)       |
| 03   | 排序+二分查找 | O(nlog(n)) | O(1)       |

```go
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	j := len(nums) - 1
	res := 0
	for i := 0; i < len(nums); i++ {
		for i < j {
			if nums[i]+nums[j] <= target {
				break
			}
			j--
		}
		if i < j {
			res = res + (j - i)
		}
	}
	return res % 1000000007
}

# 2
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	res := 0
	left, right := 0, len(nums)-1
	for left < right {
		for left < right && nums[left]+nums[right] > target {
			right--
		}
		res = res + right - left
		left++
	}
	return res % 1000000007
}

# 3
func purchasePlans(nums []int, target int) int {
	sort.Ints(nums)
	ln := len(nums)
	res := 0
	for i := 0; i < ln; i++ {
		target := target - nums[i]
		index := search(nums[i+1:], target)
		res = res + index
	}
	return res % 1000000007
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
```

## LCP29.乐团站位(2)

- 题目

```
某乐团的演出场地可视作 num * num 的二维矩阵 grid（左上角坐标为 [0,0])，每个位置站有一位成员。
乐团共有 9 种乐器，乐器编号为 1~9，每位成员持有 1 个乐器。
为保证声乐混合效果，成员站位规则为：自 grid 左上角开始顺时针螺旋形向内循环以 1，2，...，9 循环重复排列。
例如当 num = 5 时，站位如图所示
请返回位于场地坐标 [Xpos,Ypos] 的成员所持乐器编号。
示例 1：输入：num = 3, Xpos = 0, Ypos = 2 输出：3
解释：
示例 2：输入：num = 4, Xpos = 1, Ypos = 2 输出：5
解释：
提示：1 <= num <= 10^9
0 <= Xpos, Ypos < num
```

- 解题思路

| No.  | 思路   | 时间复杂度 | 空间复杂度 |
| ---- | ------ | ---------- | ---------- |
| 01   | 找规律 | O(1)       | O(1)       |
| 02   | 找规律 | O(1)       | O(1)       |

```go
func orchestraLayout(num int, xPos int, yPos int) int {
	x := min(xPos, num-1-xPos)
	y := min(yPos, num-1-yPos)
	k := min(x, y) // 在第几圈（从0开始）
	// n*n - (n-2k)*(n-2k) = n*n-(n*n+4k*k-4nk)= 4nk-4k*k
	total := 4 * k * (num - k) % 9 // 第几圈外总共的个数
	if xPos == k {                 // 上边
		return (total+yPos-k)%9 + 1
	} else if yPos == num-1-k { // 右边
		before := num - 2*k - 1
		return (total+before+xPos-k)%9 + 1
	} else if xPos == num-1-k { // 下边
		before := (num - 2*k - 1) * 2
		return (total+before+num-k-1-yPos)%9 + 1
	} else if yPos == k { // 左边
		before := (num - 2*k - 1) * 3
		return (total+before+num-k-1-xPos)%9 + 1
	}
	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 2
func orchestraLayout(num int, xPos int, yPos int) int {
	x := min(xPos, num-1-xPos)
	y := min(yPos, num-1-yPos)
	k := min(x, y) // 在第几圈（从0开始）
	// n*n - (n-2k)*(n-2k) = n*n-(n*n+4k*k-4nk)= 4nk-4k*k
	if xPos <= yPos {
		total := num*num - (num-2*k)*(num-2*k)
		return (total+xPos-k+yPos-k)%9 + 1
	} else {
		total := num*num - (num-(2*k+2))*(num-(2*k+2))
		return (total-(xPos-k)-(yPos-k))%9 + 1
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
```

## LCP30.魔塔游戏(2)

- 题目

```
小扣当前位于魔塔游戏第一层，共有 N 个房间，编号为 0 ~ N-1。
每个房间的补血道具/怪物对于血量影响记于数组 nums，其中正数表示道具补血数值，即血量增加对应数值；
负数表示怪物造成伤害值，即血量减少对应数值；0 表示房间对血量无影响。
小扣初始血量为 1，且无上限。假定小扣原计划按房间编号升序访问所有房间补血/打怪，为保证血量始终为正值，
小扣需对房间访问顺序进行调整，每次仅能将一个怪物房间（负数的房间）调整至访问顺序末尾。
请返回小扣最少需要调整几次，才能顺利访问所有房间。若调整顺序也无法访问完全部房间，请返回 -1。
示例 1：输入：nums = [100,100,100,-250,-60,-140,-50,-50,100,150] 输出：1
解释：初始血量为 1。至少需要将 nums[3] 调整至访问顺序末尾以满足要求。
示例 2：输入：nums = [-200,-300,400,0]输出：-1
解释：调整访问顺序也无法完成全部房间的访问。
提示：1 <= nums.length <= 10^5
-10^5 <= nums[i] <= 10^5
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 堆   | O(nlog(n)) | O(n)       |
| 02   | 堆   | O(nlog(n)) | O(n)       |

```go
func magicTower(nums []int) int {
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	blood := 0
	sum := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		if nums[i] < 0 {
			heap.Push(&intHeap, nums[i])
			if blood+nums[i] < 0 {
				res++
				minValue := heap.Pop(&intHeap).(int)
				blood = blood - minValue 
			}
		}
		blood = blood + nums[i]
	}
	if sum < 0 {
		return -1
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}

# 2
func magicTower(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	if sum < 0 {
		return -1
	}
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	blood := 0
	res := 0
	for i := 0; i < len(nums); i++ {
		heap.Push(&intHeap, nums[i])
		blood = blood + nums[i]
		if blood < 0 {
			minValue := heap.Pop(&intHeap).(int)
			blood = blood - minValue
			res++
		}
	}
	return res
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
```

## LCP33.蓄水(2)

- 题目

```
给定 N 个无限容量且初始均空的水缸，每个水缸配有一个水桶用来打水，
第 i 个水缸配备的水桶容量记作 bucket[i]。小扣有以下两种操作：
升级水桶：选择任意一个水桶，使其容量增加为 bucket[i]+1
蓄水：将全部水桶接满水，倒入各自对应的水缸
每个水缸对应最低蓄水量记作 vat[i]，返回小扣至少需要多少次操作可以完成所有水缸蓄水要求。
注意：实际蓄水量 达到或超过 最低蓄水量，即完成蓄水要求。
示例 1：输入：bucket = [1,3], vat = [6,8] 输出：4
解释：第 1 次操作升级 bucket[0]；
第 2 ~ 4 次操作均选择蓄水，即可完成蓄水要求。
示例 2：输入：bucket = [9,0,1], vat = [0,2,2] 输出：3
解释：第 1 次操作均选择升级 bucket[1]
第 2~3 次操作选择蓄水，即可完成蓄水要求。
提示：1 <= bucket.length == vat.length <= 100
0 <= bucket[i], vat[i] <= 10^4
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 枚举 | O(n^2)     | O(1)       |
| 02   | 堆   | O(nlog(n)) | O(n)       |

```go
func storeWater(bucket []int, vat []int) int {
	n := len(vat)
	maxValue := 0
	for i := 0; i < n; i++ {
		maxValue = max(maxValue, vat[i])
	}
	if maxValue == 0 {
		return 0
	}
	res := math.MaxInt32
	for k := 1; k <= maxValue; k++ { // 枚举蓄水的次数
		temp := k
		for i := 0; i < n; i++ {
			begin := vat[i] / k // 需要升级到的目的容量
			if vat[i]%k > 0 {
				begin++
			}
			if begin > bucket[i] {
				temp = temp + begin - bucket[i] // 升级次数
			}
		}
		res = min(res, temp)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

# 2
func storeWater(bucket []int, vat []int) int {
	n := len(vat)
	nodeHeap := make(IntHeap, 0)
	heap.Init(&nodeHeap)
	count := 0 // 需要升级的次数
	for i := 0; i < n; i++ {
		if bucket[i] == 0 && vat[i] > 0 {
			bucket[i] = 1
			count++
		}
		if vat[i] > 0 {
			heap.Push(&nodeHeap, Node{
				bucket: bucket[i],
				vat:    vat[i],
				count:  (vat[i]-1)/bucket[i] + 1,
			})
		}
	}
	res := math.MaxInt32 // 总次数
	for nodeHeap.Len() > 0 {
		node := heap.Pop(&nodeHeap).(Node)
		if count >= res {
			break
		}
        res = min(res, node.count+count) // 堆里面最大的蓄水次数+升级的次数
		heap.Push(&nodeHeap, Node{
			bucket: node.bucket + 1,
			vat:    node.vat,
			count:  (node.vat-1)/(node.bucket+1) + 1,
		})
		count++
	}
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type Node struct {
	bucket int
	vat    int
	count  int
}

type IntHeap []Node

func (h IntHeap) Len() int {
	return len(h)
}

// 小根堆<,大根堆变换方向>
func (h IntHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(Node))
}

func (h *IntHeap) Pop() interface{} {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
```

## LCP34.二叉树染色(1)

- 题目

```
小扣有一个根结点为 root 的二叉树模型，初始所有结点均为白色，可以用蓝色染料给模型结点染色，
模型的每个结点有一个 val 价值。
小扣出于美观考虑，希望最后二叉树上每个蓝色相连部分的结点个数不能超过 k 个，
求所有染成蓝色的结点价值总和最大是多少？
示例 1：输入：root = [5,2,3,4], k = 2输出：12
解释：结点 5、3、4 染成蓝色，获得最大的价值 5+3+4=12
示例 2：输入：root = [4,1,3,9,null,null,2], k = 2 输出：16
解释：结点 4、3、9 染成蓝色，获得最大的价值 4+3+9=16
提示：1 <= k <= 10
1 <= val <= 10000
1 <= 结点数量 <= 10000
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 动态规划-递归 | O(n^3)     | O(n^2)     |

```go
func maxValue(root *TreeNode, k int) int {
	dp := dfs(root, k)
	return maxArr(dp)
}

func dfs(root *TreeNode, k int) []int {
	dp := make([]int, k+1) // dp[i]表示，染色数为i的最大值
	if root == nil {
		return dp
	}
	left := dfs(root.Left, k)
	right := dfs(root.Right, k)
	dp[0] = maxArr(left) + maxArr(right) // 当前节点不染色
	for i := 1; i <= k; i++ {            // 当前节点染色
		for j := 0; j < i; j++ {
			dp[i] = max(dp[i], left[j]+right[i-1-j]+root.Val)
		}
	}
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArr(arr []int) int {
	res := 0
	for i := 0; i < len(arr); i++ {
		res = max(res, arr[i])
	}
	return res
}
```

## LCP35.电动车游城市(1)

- 题目

```
小明的电动车电量充满时可行驶距离为 cnt，每行驶 1 单位距离消耗 1 单位电量，且花费 1 单位时间。
小明想选择电动车作为代步工具。地图上共有 N 个景点，景点编号为 0 ~ N-1。
他将地图信息以 [城市 A 编号,城市 B 编号,两城市间距离] 格式整理在在二维数组 paths，表示城市 A、B 间存在双向通路。
初始状态，电动车电量为 0。每个城市都设有充电桩，charge[i] 表示第 i 个城市每充 1 单位电量需要花费的单位时间。
请返回小明最少需要花费多少单位时间从起点城市 start 抵达终点城市 end。
示例 1：输入：paths = [[1,3,3],[3,2,1],[2,1,3],[0,1,4],[3,0,5]], cnt = 6, 
start = 1, end = 0, charge = [2,10,4,1] 输出：43
解释：最佳路线为：1->3->0。
在城市 1 仅充 3 单位电至城市 3，然后在城市 3 充 5 单位电，行驶至城市 5。
充电用时共 3*10 + 5*1= 35
行驶用时 3 + 5 = 8，此时总用时最短 43。
示例 2：输入：paths = [[0,4,2],[4,3,5],[3,0,5],[0,1,5],[3,2,4],[1,2,8]], cnt = 8, 
start = 0, end = 2, charge = [4,1,1,3,2] 输出：38
解释：最佳路线为：0->4->3->2。
城市 0 充电 2 单位，行驶至城市 4 充电 8 单位，行驶至城市 3 充电 1 单位，最终行驶至城市 2。
充电用时 4*2+2*8+3*1 = 27
行驶用时 2+5+4 = 11，总用时最短 38。
提示：1 <= paths.length <= 200
paths[i].length == 3
2 <= charge.length == n <= 100
0 <= path[i][0],path[i][1],start,end < n
1 <= cnt <= 100
1 <= path[i][2] <= cnt
1 <= charge[i] <= 100
题目保证所有城市相互可以到达
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | Dijkstra | O(n^2)     | O(n^2)     |

```go
func electricCarPlan(paths [][]int, cnt int, start int, end int, charge []int) int {
	n := len(charge)
	arr := make([][][]int, n) // 邻接表
	for i := 0; i < len(paths); i++ {
		a, b, c := paths[i][0], paths[i][1], paths[i][2] // a<=>b
		arr[a] = append(arr[a], []int{b, c})
		arr[b] = append(arr[b], []int{a, c})
	}
	dis := make([][]int, n) // start到i点在j电量下的花费
	for i := 0; i < n; i++ {
		dis[i] = make([]int, cnt+1)
		for j := 0; j < cnt+1; j++ {
			dis[i][j] = math.MaxInt32
		}
	}
	dis[start][0] = 0 // 开始花费为0
	intHeap := make(IntHeap, 0)
	heap.Init(&intHeap)
	heap.Push(&intHeap, []int{0, start, 0}) // 时间+位置+电量：按时间堆排序
	for intHeap.Len() > 0 {
		node := heap.Pop(&intHeap).([]int)
		t, cur, value := node[0], node[1], node[2]
		if t > dis[cur][value] { // 大于跳过
			continue
		}
		if cur == end { // 终点，直接返回
			return t
		}
		// 核心点：可以在一个城市充满电，可以不需要在其它城市充电；
		// 这样可以尝试在一个城市充满一定电量后往下走；不一定非要完全充满电或者只充到达下一个城市所需的电量
		// 因为每个城市的充电单价不一样
		if value < cnt { // 去第cur城市充电1单位；入堆后可以继续充电1单位
			nextTime := t + charge[cur]
			if nextTime < dis[cur][value+1] {
				dis[cur][value+1] = nextTime
				heap.Push(&intHeap, []int{nextTime, cur, value + 1})
			}
		}
		// 电量满足到达下一个城市后可以开往下一个城市
		for i := 0; i < len(arr[cur]); i++ {
			next, nextDis := arr[cur][i][0], arr[cur][i][1]
			if value >= nextDis && t+nextDis < dis[next][value-nextDis] {
				dis[next][value-nextDis] = t + nextDis
				heap.Push(&intHeap, []int{t + nextDis, next, value - nextDis})
			}
		}
	}
	return -1
}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
```

## LCP36.最多牌组数

### 题目

```
麻将的游戏规则中，共有两种方式凑成「一组牌」：
顺子：三张牌面数字连续的麻将，例如 [4,5,6]
刻子：三张牌面数字相同的麻将，例如 [10,10,10]
给定若干数字作为麻将牌的数值（记作一维数组 tiles），请返回所给 tiles 最多可组成的牌组数。
注意：凑成牌组时，每张牌仅能使用一次。
示例 1：输入：tiles = [2,2,2,3,4] 输出：1
解释：最多可以组合出 [2,2,2] 或者 [2,3,4] 其中一组牌。
示例 2：输入：tiles = [2,2,2,3,4,1,3] 输出：2
解释：最多可以组合出 [1,2,3] 与 [2,3,4] 两组牌。
提示：1 <= tiles.length <= 10^5
1 <= tiles[i] <= 10^9
```

### 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n^2)     | O(n^2)     |

```go

```

## LCP39.无人机方阵(1)

- 题目

```
在 「力扣挑战赛」 开幕式的压轴节目 「无人机方阵」中，每一架无人机展示一种灯光颜色。 无人机方阵通过两种操作进行颜色图案变换：
调整无人机的位置布局
切换无人机展示的灯光颜色
给定两个大小均为 N*M 的二维数组 source 和 target 表示无人机方阵表演的两种颜色图案，
由于无人机切换灯光颜色的耗能很大，请返回从 source 到 target 最少需要多少架无人机切换灯光颜色。
注意： 调整无人机的位置布局时无人机的位置可以随意变动。
示例 1：输入：source = [[1,3],[5,4]], target = [[3,1],[6,5]] 输出：1
解释：最佳方案为
将 [0,1] 处的无人机移动至 [0,0] 处；
将 [0,0] 处的无人机移动至 [0,1] 处；
将 [1,0] 处的无人机移动至 [1,1] 处；
将 [1,1] 处的无人机移动至 [1,0] 处，其灯光颜色切换为颜色编号为 6 的灯光；
因此从source 到 target 所需要的最少灯光切换次数为 1。
示例 2：输入：source = [[1,2,3],[3,4,5]], target = [[1,3,5],[2,3,4]]输出：0
解释：仅需调整无人机的位置布局，便可完成图案切换。因此不需要无人机切换颜色
提示：n == source.length == target.length
m == source[i].length == target[i].length
1 <= n, m <=100
1 <= source[i][j], target[i][j] <=10^4
```

- 解题思路

| No.  | 思路     | 时间复杂度 | 空间复杂度 |
| ---- | -------- | ---------- | ---------- |
| 01   | 哈希辅助 | O(n^2)     | O(n^2)     |

 ```go
 func minimumSwitchingTimes(source [][]int, target [][]int) int {
 	temp := make(map[int]int)
 	n, m := len(source), len(source[0])
 	for i := 0; i < n; i++ {
 		for j := 0; j < m; j++ {
 			temp[source[i][j]]++
 			temp[target[i][j]]--
 		}
 	}
 	res := 0
 	for _, v := range temp {
 		if v > 0 {
 			res = res + v
 		}
 	}
 	return res
 }
 ```

## LCP40.心算挑战(3)

- 题目

```
「力扣挑战赛」心算项目的挑战比赛中，要求选手从 N 张卡牌中选出 cnt 张卡牌，若这 cnt 张卡牌数字总和为偶数，
则选手成绩「有效」且得分为 cnt 张卡牌数字总和。
给定数组 cards 和 cnt，其中 cards[i] 表示第 i 张卡牌上的数字。 请帮参赛选手计算最大的有效得分。
若不存在获取有效得分的卡牌方案，则返回 0。
示例 1：输入：cards = [1,2,8,9], cnt = 3 输出：18
解释：选择数字为 1、8、9 的这三张卡牌，此时可获得最大的有效得分 1+8+9=18。
示例 2：输入：cards = [3,3,1], cnt = 1 输出：0
解释：不存在获取有效得分的卡牌方案。
提示：1 <= cnt <= cards.length <= 10^5
1 <= cards[i] <= 1000
```

- 解题思路

| No.  | 思路      | 时间复杂度 | 空间复杂度 |
| ---- | --------- | ---------- | ---------- |
| 01   | 排序+枚举 | O(nlog(n)) | O(n)       |
| 02   | 排序+枚举 | O(nlog(n)) | O(n)       |
| 03   | 贪心      | O(nlog(n)) | O(1)       |

 ```go
 func maxmiumScore(cards []int, cnt int) int {
 	a, b := make([]int, 0), make([]int, 0)
 	for i := 0; i < len(cards); i++ {
 		if cards[i]%2 == 0 {
 			a = append(a, cards[i])
 		} else {
 			b = append(b, cards[i])
 		}
 	}
 	sort.Slice(a, func(i, j int) bool {
 		return a[i] > a[j]
 	})
 	sort.Slice(b, func(i, j int) bool {
 		return b[i] > b[j]
 	})
 	x, y := len(a), len(b)
 	arrA, arrB := make([]int, x+1), make([]int, y+1)
 	for i := 0; i < x; i++ {
 		arrA[i+1] = arrA[i] + a[i]
 	}
 	for i := 0; i < y; i++ {
 		arrB[i+1] = arrB[i] + b[i]
 	}
 	res := 0
 	for i := 0; i <= cnt; i++ { // 枚举奇数的个数
 		n, m := cnt-i, i
 		if n <= x && m <= y && (arrA[n]+arrB[m])%2 == 0 {
 			res = max(res, arrA[n]+arrB[m])
 		}
 	}
 	return res
 }
 
 func max(a, b int) int {
 	if a > b {
 		return a
 	}
 	return b
 }
 
 # 2
 func maxmiumScore(cards []int, cnt int) int {
 	sort.Slice(cards, func(i, j int) bool {
 		return cards[i] > cards[j]
 	})
 	a, b := make([]int, 1), make([]int, 1)
 	for i := 0; i < len(cards); i++ {
 		if cards[i]%2 == 0 {
 			a = append(a, cards[i]+a[len(a)-1])
 		} else {
 			b = append(b, cards[i]+b[len(b)-1])
 		}
 	}
 	res := 0
 	for i := 0; i <= cnt; i++ { // 枚举奇数的个数
 		n, m := cnt-i, i
 		if n < len(a) && m < len(b) && (a[n]+b[m])%2 == 0 {
 			res = max(res, a[n]+b[m])
 		}
 	}
 	return res
 }
 
 func max(a, b int) int {
 	if a > b {
 		return a
 	}
 	return b
 }
 
 # 3
 func maxmiumScore(cards []int, cnt int) int {
 	sort.Slice(cards, func(i, j int) bool {
 		return cards[i] > cards[j]
 	})
 	res := 0
 	sum := 0
 	for i := 0; i < cnt; i++ {
 		sum = sum + cards[i]
 	}
 	if sum%2 == 0 { // 偶数直接返回
 		return sum
 	}
 	// 使用不同奇偶性的数替换：
 	// 1、找个一个较小的数替换：cards[cnt-1]
 	// 2、找到一个较小的数替换：跟cards[cnt-1]不同的较小的数
 	for i := cnt; i < len(cards); i++ { // 情况1：在后面找1个数替换前cnt个最大数的最后1个数
 		if cards[i]%2 != cards[cnt-1]%2 {
 			res = max(res, sum-cards[cnt-1]+cards[i])
 			break
 		}
 	}
 	for i := cnt - 2; i >= 0; i-- { // 情况2：尝试找到1个奇偶性不同于cards[cnt-1]的数，然后替换掉
 		if cards[i]%2 != cards[cnt-1]%2 {
 			for j := cnt; j < len(cards); j++ {
 				if cards[j]%2 != cards[i]%2 {
 					res = max(res, sum-cards[i]+cards[j])
 				}
 			}
 			break
 		}
 	}
 	return res
 }
 
 func max(a, b int) int {
 	if a > b {
 		return a
 	}
 	return b
 }
 ```

## LCP41.黑白翻转棋(1)

- 题目

```
在 n*m 大小的棋盘中，有黑白两种棋子，黑棋记作字母 "X", 白棋记作字母 "O"，空余位置记作 "."。
当落下的棋子与其他相同颜色的棋子在行、列或对角线完全包围（中间不存在空白位置）另一种颜色的棋子，则可以翻转这些棋子的颜色。
「力扣挑战赛」黑白翻转棋项目中，将提供给选手一个未形成可翻转棋子的棋盘残局，其状态记作 chessboard。
若下一步可放置一枚黑棋，请问选手最多能翻转多少枚白棋。
注意：若翻转白棋成黑棋后，棋盘上仍存在可以翻转的白棋，将可以 继续 翻转白棋
输入数据保证初始棋盘状态无可以翻转的棋子且存在空余位置
示例 1：输入：chessboard = ["....X.","....X.","XOOO..","......","......"] 输出：3
解释：可以选择下在 [2,4] 处，能够翻转白方三枚棋子。
示例 2：输入：chessboard = [".X.",".O.","XO."] 输出：2
解释：可以选择下在 [2,2] 处，能够翻转白方两枚棋子。
示例 3：输入：chessboard = [".......",".......",".......","X......",".O.....","..O....","....OOX"] 输出：4
解释： 可以选择下在 [6,3] 处，能够翻转白方四枚棋子。
提示：1 <= chessboard.length, chessboard[i].length <= 8
chessboard[i] 仅包含 "."、"O" 和 "X"
```

- 解题思路

| No.  | 思路              | 时间复杂度 | 空间复杂度 |
| ---- | ----------------- | ---------- | ---------- |
| 01   | 枚举+深度优先搜索 | O(n^4)     | O(n^2)     |

```go
var res int
var n, m int
var dx = []int{-1, 1, 0, 0, 1, 1, -1, -1}
var dy = []int{0, 0, -1, 1, 1, -1, -1, 1}

func flipChess(chessboard []string) int {
	res = 0
	n, m = len(chessboard), len(chessboard[0])
	temp := make([][]byte, n)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if chessboard[i][j] == '.' {
				for a := 0; a < n; a++ {
					temp[a] = []byte(chessboard[a])
				}
				temp[i][j] = 'X'
				dfs(temp, i, j) // 尝试在i,j位置下黑棋
				count := 0
				for a := 0; a < n; a++ { // 统计结果
					for b := 0; b < m; b++ {
						if temp[a][b] == 'X' && chessboard[a][b] == 'O' {
							count++
						}
					}
				}
				res = max(res, count) // 更新结果
			}
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func dfs(arr [][]byte, i, j int) {
	arr[i][j] = 'X'
	for k := 0; k < 8; k++ {
		if ok, path := judge(arr, i, j, 'X', dx[k], dy[k]); ok == true {
			for c := 0; c < len(path); c++ {
				a, b := path[c][0], path[c][1]
				arr[a][b] = 'X'
				dfs(arr, a, b)
			}
		}
	}
}

// leetcode1958.检查操作是否合法
func judge(board [][]byte, rMove int, cMove int, color byte, dirX, dirY int) (res bool, path [][]int) {
	x, y := rMove+dirX, cMove+dirY
	count := 1
	for 0 <= x && x < n && 0 <= y && y < m {
		if board[x][y] == '.' {
			return false, nil
		}
		path = append(path, []int{x, y})
		if count == 1 {
			if board[x][y] == color {
				return false, nil
			}
		} else {
			if board[x][y] == color {
				return true, path
			}
		}
		count++
		x = x + dirX
		y = y + dirY
	}
	return false, nil
}
```

## LCP42.玩具套圈

### 题目

```

```

### 解题思路

```go
```

## LCP44.开幕式焰火(1)

- 题目

```
「力扣挑战赛」开幕式开始了，空中绽放了一颗二叉树形的巨型焰火。
给定一棵二叉树 root 代表焰火，节点值表示巨型焰火这一位置的颜色种类。请帮小扣计算巨型焰火有多少种不同的颜色。
示例 1：输入：root = [1,3,2,1,null,2] 输出：3
解释：焰火中有 3 个不同的颜色，值分别为 1、2、3
示例 2：输入：root = [3,3,3] 输出：1
解释：焰火中仅出现 1 个颜色，值为 3
提示：1 <= 节点个数 <= 1000
1 <= Node.val <= 1000
```

- 解题思路

| No.  | 思路 | 时间复杂度 | 空间复杂度 |
| ---- | ---- | ---------- | ---------- |
| 01   | 哈希 | O(n)       | O(n)       |

```go
var m map[int]int

func numColor(root *TreeNode) int {
	m = make(map[int]int)
	dfs(root)
	return len(m)
}

func dfs(root *TreeNode) {
	if root != nil {
		m[root.Val] = 1
		dfs(root.Left)
		dfs(root.Right)
	}
}
```

## LCP45.自行车炫技赛场(2)

- 题目

```
「力扣挑战赛」中 N*M 大小的自行车炫技赛场的场地由一片连绵起伏的上下坡组成，
场地的高度值记录于二维数组 terrain 中，场地的减速值记录于二维数组 obstacle 中。
若选手骑着自行车从高度为 h1 且减速值为 o1 的位置到高度为 h2 且减速值为 o2 的相邻位置（上下左右四个方向），
速度变化值为 h1-h2-o2（负值减速，正值增速）。
选手初始位于坐标 position 处且初始速度为 1，请问选手可以刚好到其他哪些位置时速度依旧为 1。
请以二维数组形式返回这些位置。若有多个位置则按行坐标升序排列，若有多个位置行坐标相同则按列坐标升序排列。
注意： 骑行过程中速度不能为零或负值
示例 1：输入：position = [0,0], terrain = [[0,0],[0,0]], obstacle = [[0,0],[0,0]] 输出：[[0,1],[1,0],[1,1]]
解释：由于当前场地属于平地，根据上面的规则，选手从[0,0]的位置出发都能刚好在其他处的位置速度为 1。
示例 2：输入：position = [1,1], terrain = [[5,0],[0,6]], obstacle = [[0,6],[7,0]] 输出：[[0,1]]
解释：选手从 [1,1] 处的位置出发，到 [0,1] 处的位置时恰好速度为 1。
提示：n == terrain.length == obstacle.length
m == terrain[i].length == obstacle[i].length
1 <= n <= 100
1 <= m <= 100
0 <= terrain[i][j], obstacle[i][j] <= 100
position.length == 2
0 <= position[0] < n
0 <= position[1] < m
```

- 解题思路

| No.  | 思路         | 时间复杂度 | 空间复杂度 |
| ---- | ------------ | ---------- | ---------- |
| 01   | 深度优先搜索 | O(n^2)     | O(n^2)     |
| 02   | 广度优先搜索 | O(n^2)     | O(n^2)     |

```go
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}
var res [][]int
var visited map[[3]int]bool
var n, m int
var originX, originY int

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	res = make([][]int, 0)
	n, m = len(terrain), len(terrain[0])
	originX, originY = position[0], position[1]
	visited = make(map[[3]int]bool)
	visited[[3]int{originX, originY, 1}] = true
	dfs(terrain, obstacle, originX, originY, 1)
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})
	return res
}

func dfs(terrain [][]int, obstacle [][]int, i, j int, speed int) {
	if speed == 1 && (i == originX && j == originY) == false {
		res = append(res, []int{i, j})
	}
	for k := 0; k < 4; k++ {
		x, y := i+dx[k], j+dy[k]
		if 0 <= x && x < n && 0 <= y && y < m {
			// next = speed + h1-h2-o2
			next := speed + (terrain[i][j] - terrain[x][y] - obstacle[x][y])
			if next > 0 && visited[[3]int{x, y, next}] == false {
				visited[[3]int{x, y, next}] = true
				dfs(terrain, obstacle, x, y, next)
			}
		}
	}
}

# 2
var dx = []int{0, 1, 0, -1}
var dy = []int{1, 0, -1, 0}

func bicycleYard(position []int, terrain [][]int, obstacle [][]int) [][]int {
	res := make([][]int, 0)
	n, m := len(terrain), len(terrain[0])
	originX, originY := position[0], position[1]
	visited := make(map[[3]int]bool)
	visited[[3]int{originX, originY, 1}] = true
	queue := make([][3]int, 0)
	queue = append(queue, [3]int{originX, originY, 1})
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		i, j, speed := node[0], node[1], node[2]
		if speed == 1 && (i == originX && j == originY) == false {
			res = append(res, []int{i, j})
		}
		for k := 0; k < 4; k++ {
			x, y := i+dx[k], j+dy[k]
			if 0 <= x && x < n && 0 <= y && y < m {
				// next = speed + h1-h2-o2
				next := speed + (terrain[i][j] - terrain[x][y] - obstacle[x][y])
				if next > 0 && visited[[3]int{x, y, next}] == false {
					visited[[3]int{x, y, next}] = true
					queue = append(queue, [3]int{x, y, next})
				}
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i][0] == res[j][0] {
			return res[i][1] < res[j][1]
		}
		return res[i][0] < res[j][0]
	})
	return res
}
```

## LCP46.志愿者调配(1)

- 题目

```
「力扣挑战赛」有 n 个比赛场馆（场馆编号从 0 开始），场馆之间的通道分布情况记录于二维数组 edges 中，
edges[i]= [x, y] 表示第 i 条通道连接场馆 x 和场馆 y(即两个场馆相邻)。
初始每个场馆中都有一定人数的志愿者（不同场馆人数可能不同），后续 m 天每天均会根据赛事热度进行志愿者人数调配。
调配方案分为如下三种：
将编号为 idx 的场馆内的志愿者人数减半；
将编号为 idx 的场馆相邻的场馆的志愿者人数都加上编号为 idx 的场馆的志愿者人数；
将编号为 idx 的场馆相邻的场馆的志愿者人数都减去编号为 idx 的场馆的志愿者人数。
所有的调配信息记录于数组 plans 中，plans[i] = [num,idx] 表示第 i 天对编号 idx 的场馆执行了第 num 种调配方案。
在比赛结束后对调配方案进行复盘时，不慎将第 0 个场馆的最终志愿者人数丢失，只保留了初始所有场馆的志愿者总人数 totalNum ，
以及记录了第 1 ~ n-1 个场馆的最终志愿者人数的一维数组 finalCnt。
请你根据现有的信息求出初始每个场馆的志愿者人数，并按场馆编号顺序返回志愿者人数列表。
注意：测试数据保证当某场馆进行第一种调配时，该场馆的志愿者人数一定为偶数；
测试数据保证当某场馆进行第三种调配时，该场馆的相邻场馆志愿者人数不为负数；
测试数据保证比赛开始时每个场馆的志愿者人数都不超过 10^9；
测试数据保证给定的场馆间的道路分布情况中不会出现自环、重边的情况。
示例 1：输入：finalCnt = [1,16], totalNum = 21, edges = [[0,1],[1,2]], 
plans = [[2,1],[1,0],[3,0]] 输出：[5,7,9]
解释：
示例 2 ：输入：finalCnt = [4,13,4,3,8], totalNum = 54, edges = [[0,3],[1,3],[4,3],[2,3],[2,5]], 
plans = [[1,1],[3,3],[2,5],[1,0]] 输出：[10,16,9,4,7,8]
提示：2 <= n <= 5*10^4
1 <= edges.length <= min((n * (n - 1)) / 2, 5*10^4)
0 <= edges[i][0], edges[i][1] < n
1 <= plans.length <= 10
1 <= plans[i][0] <=3
0 <= plans[i][1] < n
finalCnt.length = n-1
0 <= finalCnt[i] < 10^9
0 <= totalNum < 5*10^13
```

- 解题思路

| No.  | 思路        | 时间复杂度 | 空间复杂度 |
| ---- | ----------- | ---------- | ---------- |
| 01   | 模拟-解方程 | O(n)       | O(n)       |

```go
type Node struct {
	a float64 // 未知final[0] x的系数
	b int     // 已知的系数
}

func volunteerDeployment(finalCnt []int, totalNum int64, edges [][]int, plans [][]int) []int {
	n := len(finalCnt) + 1
	arr := make([][]int, n) // 邻接表
	for i := 0; i < len(edges); i++ {
		a, b := edges[i][0], edges[i][1]
		arr[a] = append(arr[a], b)
		arr[b] = append(arr[b], a)
	}
	final := make([]Node, 0)
	final = append(final, Node{a: 1, b: 0}) // final为最终的结果，a为final[0]未知x的系数
	for i := 0; i < len(finalCnt); i++ {
		final = append(final, Node{a: 0, b: finalCnt[i]})
	}
	for i := len(plans) - 1; i >= 0; i-- { // 倒推
		a, b := plans[i][0], plans[i][1]
		if a == 1 {
			final[b] = Node{
				a: final[b].a * 2,
				b: final[b].b * 2,
			}
		} else if a == 2 {
			for j := 0; j < len(arr[b]); j++ {
				next := arr[b][j]
				final[next] = Node{
					a: final[next].a - final[b].a,
					b: final[next].b - final[b].b,
				}
			}
		} else if a == 3 {
			for j := 0; j < len(arr[b]); j++ {
				next := arr[b][j]
				final[next] = Node{
					a: final[next].a + final[b].a,
					b: final[next].b + final[b].b,
				}
			}
		}
	}
	a, b := float64(0), int64(0)
	for i := 0; i < len(final); i++ {
		a = a + final[i].a
		b = b + int64(final[i].b)
	}
	per := float64(totalNum-b) / a // 求x的值
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = int(final[i].a*per) + final[i].b
	}
	return res
}
```

## LCP47.入场安检(2)

- 题目

```
「力扣挑战赛」 的入场仪式马上就要开始了，由于安保工作的需要，
设置了可容纳人数总和为 M 的 N 个安检室，capacities[i] 记录第 i 个安检室可容纳人数。安检室拥有两种类型：
先进先出：在安检室中的所有观众中，最早进入安检室的观众最先离开
后进先出：在安检室中的所有观众中，最晚进入安检室的观众最先离开
恰好 M+1 位入场的观众（编号从 0 开始）需要排队依次入场安检， 入场安检的规则如下：
观众需要先进入编号 0 的安检室
当观众将进入编号 i 的安检室时（0 <= i < N)，
若安检室未到达可容纳人数上限，该观众可直接进入；
若安检室已到达可容纳人数上限，在该观众进入安检室之前需根据当前安检室类型选择一位观众离开后才能进入；
当观众离开编号 i 的安检室时 （0 <= i < N-1)，将进入编号 i+1 的安检室接受安检。
若可以任意设定每个安检室的类型，请问有多少种设定安检室类型的方案可以使得编号 k 的观众第一个通过最后一个安检室入场。
注意：观众不可主动离开安检室，只有当安检室容纳人数达到上限，且又有新观众需要进入时，才可根据安检室的类型选择一位观众离开；
由于方案数可能过大，请将答案对 1000000007 取模后返回。
示例 1：输入：capacities = [2,2,3], k = 2 输出：2
解释：存在两种设定的 2 种方案：
方案 1：将编号为 0 、1 的实验室设置为 后进先出 的类型，编号为 2 的实验室设置为 先进先出 的类型；
方案 2：将编号为 0 、1 的实验室设置为 先进先出 的类型，编号为 2 的实验室设置为 后进先出 的类型。
以下是方案 1 的示意图：
示例 2：输入：capacities = [3,3], k = 3 输出：0
示例 3：输入：capacities = [4,3,2,2], k = 6 输出：2
提示:1 <= capacities.length <= 200
1 <= capacities[i] <= 200
0 <= k <= sum(capacities)
```

- 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 动态规划-二维 | O(n^2)     | O(n^2)     |
| 02   | 动态规划-一维 | O(n^2)     | O(n)       |

```go
var mod = 1000000007

func securityCheck(capacities []int, k int) int {
	n := len(capacities)
	dp := make([][]int, n+1) // dp[i][j] => 使得i个实验室的情况下第j个人 第1个离开的方案数量
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		// 先进先出不改变 入场顺序
		// 改变为 后进先出 的入场顺序：会留住 capacities[i]-1 个人
		// 把c（capacities[i]-1）看成物品大小, k看成背包容量，等价求 01背包 的方案数。
		// 从后向前遍历
		c := capacities[i-1] - 1
		for j := 0; j <= k; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= c {
				dp[i][j] = (dp[i][j] + dp[i-1][j-c]) % mod
			}
		}
	}
	return dp[n][k]
}

# 2
var mod = 1000000007

func securityCheck(capacities []int, k int) int {
	dp := make([]int, k+1) // dp[i] => 使得第i个人 第1个离开的方案数量
	dp[0] = 1
	for i := 0; i < len(capacities); i++ {
		// 先进先出不改变 入场顺序
		// 改变为 后进先出 的入场顺序：会留住 capacities[i]-1 个人
		// 把c（capacities[i]-1）看成物品大小, k看成背包容量，等价求 01背包 的方案数。
		// 从后向前遍历
		c := capacities[i] - 1
		for j := k; j >= c; j-- {
			dp[j] = (dp[j] + dp[j-c]) % mod
		}
	}
	return dp[k]
}
```

