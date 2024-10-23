package solution;

import java.util.Set;
import java.util.HashSet;

public class Solution {
    private final static int WIDTH = 9;
    private final static int HEIGHT = 9;
    public static boolean validSudoku(char[][] board) {
        for (int i=0;i < 9; i++) {
            final Set<Character> horizontalChars = new HashSet<>();
            for (int w = 0; w < WIDTH; w++) {
                final char c = board[i][w];
                if (c != '.' && horizontalChars.contains(c)) {
                    return false;
                }
                horizontalChars.add(c);
            }

            final Set<Character> verticalChars = new HashSet<>();
            for (int h = 0; h < HEIGHT; h++) {
                final char c = board[h][i];
                if (c != '.' && verticalChars.contains(c)) {
                    return false;
                }
                verticalChars.add(c);
            }
        }
        for (int i=0; i < WIDTH; i=i+3) {
            for (int j=0; j < HEIGHT; j=j+3) {
                if (!check3x3Block(board, i, j)) {
                    return false;
                }
            }
        }
        return true;
    }

    private static boolean check3x3Block(char[][] board, int x, int y) {
        if (x > WIDTH - 3 || y > HEIGHT - 3 || x < 0 || y < 0) {
            throw new IllegalArgumentException("Invalid x/y coordinates provided to check3x3Block");
        }
        final Set<Character> blockChars = new HashSet<>();
        for (int i=x;i<x+3;i++) {
            for (int j=y;j<y+3;j++) {
                final char c = board[i][j];
                if (c != '.' && blockChars.contains(c)) {
                    return false;
                }
                blockChars.add(c);
            }
        }
        return true;
    }
}