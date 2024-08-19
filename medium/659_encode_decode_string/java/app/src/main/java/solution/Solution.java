package solution;

import java.util.Map;
import java.util.HashMap;
import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;

public class Solution {
    public static List<String> encodeDecode(List<String> strs) {
        return decode(encode(strs));
    }

    private static String encode(List<String> strs) {
        final StringBuilder builder = new StringBuilder();
        for (String str : strs) {
            builder.append(str.length()).append("#").append(str);
        }
        return builder.toString();
    }

    // 4#neet4#code
    private static List<String> decode(String str) {
        final List<String> res = new ArrayList();
        int i = 0;
        while (i < str.length()) {
            int j = i;
            while (str.charAt(j) != '#') {
                j++;
            }
            int length = Integer.parseInt(str.substring(i, j));
            int wordStart = j + 1;
            res.add(str.substring(wordStart, wordStart + length));
            i = wordStart + length;
        }

        return res;
    }
}