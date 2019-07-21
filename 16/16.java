import java.util.Arrays;

class Solution {
    public int threeSumClosest(int[] nums, int target) {
        Arrays.sort(nums);
        int len = nums.length;
        int ans = nums[0] + nums[1] + nums[2];
        for (int i = 0; i < len - 2; i++) {
            int l = i + 1;
            int r = len - 1;

            while (l < r) {
                int tmp = nums[i] + nums[l] + nums[r];
                if (Math.abs(target - tmp) < Math.abs(target - ans)) {
                    ans = tmp;
                }
                if (ans == target) {
                    return ans;
                }
                if (tmp > target) {
                    r--;
                } else {
                    l++;
                }
            }
        }
        return ans;
    }
}