class Solution {
    public int maxRepOpt1(String text) {
        int[] left = new int[text.length()];
        int[] right = new int[text.length()];
        int[] num = new int[26];
        for (int i = 0; i < text.length(); i++) {
            if (i > 0 && text.charAt(i - 1) == text.charAt(i)) {
                left[i] = left[i - 1] + 1;
            } else {
                left[i] = 1;
            }
            num[(int) (text.charAt(i) - 'a')]++;
        }

        for (int i = text.length() - 1; i >= 0; i--) {
            if (i + 1 < text.length() && text.charAt(i + 1) == text.charAt(i)) {
                right[i] = right[i + 1] + 1;
            } else {
                right[i] = 1;
            }
        }

        int ans = 0;
        for (int i = 0; i < text.length(); i++) {
            int tmp = left[i];
            if (num[(int) (text.charAt(i) - 'a')] > tmp) {
                tmp++;
            }
            ans = Math.max(ans, tmp);
        }
        for (int i = 1; i < text.length() - 1; i++) {
            if (text.charAt(i - 1) == text.charAt(i + 1)) {
                int tmp = left[i - 1] + right[i + 1];
                if (num[(int) (text.charAt(i - 1) - 'a')] > tmp) {
                    tmp++;
                }
                ans = Math.max(ans, tmp);
            }
        }
        return ans;
    }
}

