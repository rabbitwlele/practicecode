class Solution {

    public int longestDecomposition(String text) {
        int ans = 0;
        int start = 0;
        int end = start + 1;
        while (true) {
            ans++;
            for (; ; end++) {
                int start1 = text.length() - end;
                int end1 = text.length() - start;
                System.out.printf("int-----  %d %d %d %d\n",start,end,start1,end1);
                if (end * 2 > text.length()) {
                    return ans;
                }
                System.out.printf("%s %s\n",text.substring(start, end),text.substring(start1, end1));
                if (text.substring(start, end).equals(text.substring(start1, end1))) {
                    System.out.printf("%d %d %d %d\n",start,end,start,end);
                    ans++;
                    break;
                }
            }
            start = end;
            end = start + 1;
        }
    }

    public static void main(String[] args) {
        Solution s = new Solution();
        System.out.println(s.longestDecomposition("aaxaa"));
    }

}
