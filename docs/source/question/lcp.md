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

## LCP19.秋叶收藏集

### 题目

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

### 解题思路

| No.  | 思路          | 时间复杂度 | 空间复杂度 |
| ---- | ------------- | ---------- | ---------- |
| 01   | 动态规划-二维 | O(n)       | O(n)       |
| 02   | 动态规划-一维 | O(n)       | O(1)       |

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
			dp[i][1] = dp[i-1][0] + 1 // 全r+当前r，需要改变一个y，步数+1
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1]+1)
				dp[i][2] = dp[i-1][1]
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2])
			}
		} else {
			dp[i][0] = dp[i-1][0] + 1 // 需要改变，步数+1
			dp[i][1] = dp[i-1][0]     // 前一个全r+当前y,不需要改变
			if i > 1 {
				dp[i][1] = min(dp[i][1], dp[i-1][1])
				dp[i][2] = dp[i-1][1] + 1
			}
			if i > 2 {
				dp[i][2] = min(dp[i][2], dp[i-1][2]+1)
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

