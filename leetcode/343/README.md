# Integer Break
[数字切分]https://leetcode.com/problems/integer-break/
> 将一个数字根据综合楼最少分成两份，是每分的乘积最大
## 分析
属于水DP吧
问题就在于最少为2份，我这里采用的方法是，初始化的时候，只从1～n-1。从这些进行衍生，这样避免了只有一次的情况