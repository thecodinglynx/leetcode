package solution;

public class ContainerWithMostWater {
    public int maxArea(int[] height) {
        var res = 0;
        int l = 0;
        int r = height.length - 1;
        while (l < r) {
            int cur = Math.min(height[l], height[r]) * (r - l);
            if (cur > res) {
                res = cur;
            }
            if (height[l] < height[r]) {
                l++;
            } else {
                r--;
            }
        }
        return res;
    }
}