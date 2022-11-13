package solution;

public class ReverseString {
    public void reverseString(char[] s) {
        int p1 = 0;
        int p2 = s.length - 1;
        while (p1 < p2) {
            char tmp;
            tmp = s[p1];
            s[p1] = s[p2];
            s[p2] = tmp;
            p1++;
            p2--;
        }
    }
}