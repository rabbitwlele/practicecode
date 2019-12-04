import java.util.ArrayList;
import java.util.List;

class Solution {
    public int[] gardenNoAdj(int N, int[][] paths) {
        int[] colors = new int[N];
        List<Integer>[] map = new List[N];
        for (int i = 0; i < N; i++) {
            map[i] = new ArrayList<Integer>();
        }

        for (int i = 0; i < paths.length; i++) {
            map[paths[i][0] - 1].add(paths[i][1] - 1);
            map[paths[i][1] - 1].add(paths[i][0] - 1);
        }

        for (int i = 0; i < N; i++) {
            int mask = 0;

            for (Integer node : map[i]) {
                mask |= (1 << colors[node.intValue()]);
            }

            for (int j = 1; j <= 4; j++) {
                if ((mask & (1 << j)) == 0) {
                    colors[i] = j;
                    break;
                }
            }
        }
        for (int x : colors){
            System.out.println(x);
        }
        return colors;
    }

}