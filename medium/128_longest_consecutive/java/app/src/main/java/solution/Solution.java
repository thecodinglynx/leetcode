package main.java.solution;

import java.util.Arrays;
import java.util.HashSet;
import java.util.Set;

public class Solution {
    public static int longestConsecutiveNLogN(int[] nums) {
        if (nums.length <= 0) {
            return 0;
        }

        Arrays.sort(nums);

        int longest = 1;
        int curLongest = 1;
        for (int i = 1; i < nums.length; i++) {
            if (nums[i] == nums[i-1]) {
                continue;
            }
            if (nums[i] - nums[i-1] == 1) {
                curLongest++;
            } else {
                if (curLongest > longest) {
                    longest = curLongest;
                }
                curLongest = 1;
            }
        }

        return Math.max(longest, curLongest);
    }

    public static int longestConsecutiveN(int[] nums) {
        if (nums.length <= 0) {
            return 0;
        }

        final var numsAsSet = new HashSet<>();
        for (int nr : nums) {
            numsAsSet.add(nr);
        }

        int longest = 1;
        int curLongest = 1;
        for (int nr : nums) {
            if (!numsAsSet.contains(nr-1)) {
                // this is a start number
                curLongest = 1;
                int nextNr = nr;
                while (true) {
                    nextNr++;
                    if (numsAsSet.contains(nextNr)) {
                        curLongest++;
                    } else {
                        if (curLongest > longest) {
                            longest = curLongest;
                        }
                        break;
                    }
                }
            }
        }

        return Math.max(longest, curLongest);
    }
}