class Solution(object):
    def maxCount(self, m, n, ops):
        """
        :type m: int
        :type n: int
        :type ops: List[List[int]]
        :rtype: int
        """
        mMin = m
        nMin = n
        for i in range(len(ops)):
            mMin = min(mMin,ops[i][0])
            nMin = min(nMin,ops[i][1])
        return nMin*mMin