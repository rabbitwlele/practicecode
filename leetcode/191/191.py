class Solution(object):
    def hammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        ans=0
        while(n):
            ans += n&1
            n>>=1
        return ans
    def beterHammingWeight(self, n):
        """
        :type n: int
        :rtype: int
        """
        count = 0
        while n > 0:
            count += 1
            n &= (n - 1)
        return count
print Solution().hammingWeight(1)