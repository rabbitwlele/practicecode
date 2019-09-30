class Solution {
    private  final int mod = 1000000000 + 7;

    public int numRollsToTarget(int d, int f, int target) {
        int[][] dp = new int[2][1001];
        int now = 1;
        dp[0][0] = 1;
        for( int i = 0 ; i < d ; i++){
            for( int j = 0 ; j <= target ;j++)
                dp[now][j] = 0;
            for ( int j = 1 ; j<=f ; j++){
                for( int from = 0 ; from + j<= target ; from++)
                dp[now][from+j] = (dp[now][from+j]+dp[now^1][from])%mod;
            }
            now ^=1;
        }
        return dp[now^1][target];
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.numRollsToTarget(30,30,500));
        System.out.println(s.numRollsToTarget(2,5,10));
    }
}