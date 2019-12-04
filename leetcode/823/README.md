#  Binary Trees With Factors
[乘积二叉树](https://leetcode.com/problems/binary-trees-with-factors/)
> 给定一个不存在重复数字的数组，可以重复使用里面的数字。我们定义一个二叉树的结构为，对于非叶子节点的只等于左右节点的值的乘积。求这样的二叉树有多少个。
## 分析
这种结果个数特别多的题目，一般都是递推公式。

我们设cnt[k]以k节点为根节点的树的个数。对于任意cnt[k]，对于所有的 i，j 如果i*j==k，cnt[k] 等于所有sum(cnt[i]*cnt[j])。

对于递推公式很重要的一点，就是执行顺序，我们在枚举i，j的时候，必须要保证i，j已经计算完毕。

所以我思考的如下条件，i，j始终满足j<=i，这样可以保证我执行到i的时候，cnt[i]已经计算完毕。由于i，j可以交换，所以我们做了二倍的处理。
```go
for i := 0; i < len(A); i++ {
	for j := 0; j <= i; j++ {
		to := A[i] * A[j]
		if cnt, ok := cnts[to]; ok {
			if i == j {
				cnts[to] = (cnt + (cnts[A[i]]*cnts[A[j]])%mod) % mod
			} else {
				cnts[to] = (cnt + (cnts[A[i]]*cnts[A[j]]<<1)%mod) % mod
			}
		}
	}
}
```