package solution;

public class Solution {

    // prefix  (init = 1):
    // num:     1 | 2 | 4 | 6
    // res:     1 | 1 | 2 | 8
    // postfix (init = 1):
    // num:     1 | 2 | 4 | 6
    // res:    48 |24 |12 | 8
    public static int[] productExceptSelf(int[] nums) {
        int[] res = new int[nums.length];

        int prefix = 1;
        for (int i = 0; i < nums.length; i++) {
            res[i] = prefix;
            prefix *= nums[i];
        }
        int postfix = 1;
        for (int i = nums.length - 1; i >= 0; i--) {
            res[i] *= postfix;
            postfix *= nums[i];
        }
        return res;
    }
}