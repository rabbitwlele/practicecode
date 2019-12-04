class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        flag = x<0
        ans = 0
        if x<0 :
            x = -x
        while(x!=0):
            ans *= 10
            ans += x%10
            x /=10
        if flag:
            ans = -ans
        if ans > 2**31-1 or ans < -2**31:
            ans = 0
        return ans 
print 2**31-1,-2**31
print Solution().reverse(10)
print Solution().reverse(1230)
print Solution().reverse(12003)
print Solution().reverse(1534236469)
print Solution().reverse(-10)
print Solution().reverse(-1230)
print Solution().reverse(-12003)