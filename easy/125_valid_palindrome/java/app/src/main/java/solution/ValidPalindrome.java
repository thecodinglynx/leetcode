package main.java.solution;

public class ValidPalindrome {
    public static boolean validate(final String str) {
        final var sanitized = str.replaceAll("[^a-zA-Z0-9]", "").toLowerCase();
        System.out.println(sanitized);
        for (int i = 0; i < sanitized.length()/2; i++) {
            if (sanitized.charAt(i) != sanitized.charAt(sanitized.length()-1-i)) {
                return false;
            }
        }
        return true;
    }
}