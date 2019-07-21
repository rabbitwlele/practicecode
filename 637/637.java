import java.util.Queue;

/**
 * Definition for a binary tree node. public class TreeNode { int val; TreeNode
 * left; TreeNode right; TreeNode(int x) { val = x; } }
 */
public class TreeNode {
    int val;
    TreeNode left;
    TreeNode right;

    TreeNode(int x) {
        val = x;
    }
}

class Solution {
    public List<Double> averageOfLevels(TreeNode root) {
        Queue<TreeNode> que = new LinkedList<TreeNode>();
        List<Double> ans = new ArrayList<Double>();
        que.offer(root);
        while (!que.isEmpty()) {
            int size = que.size();
            int tmp = size;
            long sum = 0;
            while (tmp!=0) {
                tmp--;
                TreeNode node = que.poll();
                sum +=node.val;
                if(node.left !=null){
                    que.offer(node.left);
                }
                if (node.right !=null){
                    que.offer(node.right);
                }
            }
            ans.add(new Double((double)sum/(double)size));
        }
        return ans;
    }
}