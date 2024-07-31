package solution;

public class Solution {
    public static boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        int[] characters = new int[26];
        for (int i=0;i<s.length();i++) {
            characters[s.charAt(i) - 'a']++;
            characters[t.charAt(i) - 'a']--;
        }

        for (int c : characters) {
            if (c != 0) {
                return false;
            }
        }
        return true;
    }
}