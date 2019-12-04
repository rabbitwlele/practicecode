# Combination Sum II
[组合的个数](https://leetcode.com/problems/combination-sum-ii/)
> 给一个数组，存在重复的数字，找到所有组合的和等于target。组合之前不能相等
## 分析
我们吸取了[47题](/leetcode/47)的经验，还是把相同的数字看作一个整体，以关心个数，所以第一层是枚举第几个，第二层是枚举的使用多少个。