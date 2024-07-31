package solution;

import java.util.Map;
import java.util.HashMap;

public class Solution {
    public static boolean hasDuplicate(int[] nums) {
        Map<Integer, Integer> dups = new HashMap<>();
        for (int num : nums) {
            if (dups.containsKey(num)) {
                return true;
            }
            dups.put(num, 1);
        }
        return false;
    }
}