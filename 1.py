class Solution(object):
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """
        m = {}
        for idx,i in enumerate(nums):
            if i in m:
                m[i].append(idx)
            else:
                m[i]=[idx]
        for i in nums:
            j = target-i
            if j!=i:
                if j in m:
                    return [m[i][0],m[j][0]]
            else:
                if j in m and len(m[j])>=2:
                    return m[j][:2]

print Solution().twoSum([1,1,1,4],2)