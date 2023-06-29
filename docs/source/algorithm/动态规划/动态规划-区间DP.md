# 动态规划-区间动态规划（区间DP）

- 参考：

## 0、定义

## 1、操作

## 2、Go实现

```go
for(int len = 1; len <= n; len++)          //区间长度
    for(int i = 1; i + len - 1 <= n; i++)  //枚举起点
    {
        int j = i + len - 1;               //区间终点
        // if(len == 1) dp[i][j] 初始化
        for(int k = i; k < j; k++)         //枚举分割点，构造状态转移方程
            dp[i][j] = max(dp[i][j], dp[i][k] + dp[k+1][j] + cost);
        
    }
```

## 3、Leetcode

| Title                                                                                                            | Tag         | 难度     | 完成情况 |
| :------------------------------------------------------------------------------------------------------------------| :-------------| :--------| :------|
| [312.戳气球](https://leetcode.cn/problems/burst-balloons/)                                                      | 分治算法、动态规划   | Hard   | 完成   |
| [664.奇怪的打印机](https://leetcode.cn/problems/strange-printer/)                                                  | 深度优先搜索、动态规划 | Hard   | 完成   |
| [1039.多边形三角剖分的最低得分](https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/)                    | 动态规划        | Medium | 完成   |
| [1547.切棍子的最小成本](https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/)                                   | 动态规划        | Hard   | 完成   |
| [1959.K次调整数组大小浪费的最小总空间](https://leetcode.cn/problems/minimum-total-space-wasted-with-k-resizing-operations/) | 数组、动态规划     | Medium | 完成   |