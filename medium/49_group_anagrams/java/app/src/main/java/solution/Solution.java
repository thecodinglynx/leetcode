package solution;

import java.util.Map;
import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    public static List<List<String>> groupAnagrams(String[] strs) {
        final Map<String, List<String>> res = new HashMap<>();

        for (String s : strs) {
            int[] chars = new int[26];
            for (char c : s.toCharArray()) {
                chars[c - 'a']++;
            }
            final String key = Arrays.toString(chars);
            if (!res.containsKey(key)) {
                res.put(key, new ArrayList<>());
            }
            res.get(key).add(s);
        }
        return new ArrayList<>(res.values());
    }
}