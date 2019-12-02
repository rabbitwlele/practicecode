# Palindrome Partitioning III
[分隔会文串](https://leetcode.com/problems/palindrome-partitioning-iii/)
> 给你一个字符串，将该字符串分成k份，可以将任意一个字符串替换，保证每一份都是回文串。求最小的替换次数

## 分析
这个题目不算困难，DP的味道非常明显。

设字符串为s

dp[i][k]代表 s[0,i)，分隔为k份，最小的替换次数
```go
for m := 0; m < i; m++ {
	dp[i][j] = min(dp[i][j], dp[m][j-1]+ 替换次数(s[m:j]))
}
```