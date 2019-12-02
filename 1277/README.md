# Count Square Submatrices with All Ones
[Count Square Submatrices with All Ones](https://leetcode.com/problems/count-square-submatrices-with-all-ones/)
> Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.
## 题目
**只包含1的子正方形的个数**
> 给出一个m*n的，并且只有1和0的矩阵，返回只包含1的正方形的个数
## 分析
对与一个长度为n的正方形，改正方形一个共有n个子正方形。

dp[i][j] 代表以i，j作为右下角的，并且只包含1的正方形的最长的长度。

所有dp[i][j] 就是只包含1的子正方形的个数。

问题就转换为如何求出dp[i][j]，我们先给出一个递推方程

```go
if dp[i][j] == 1 {
    dp[i][j] += min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1]))
}
```

在dp[i][j]为1的情况下，dp[i][j]就等于dp[i-1][j-1],dp[i-1][j],dp[i][j-1]中最小值+1

本人水平优先，无法从数学上给出证明，所以上一个图比较直接

![](/1277/pics/1277.png)

