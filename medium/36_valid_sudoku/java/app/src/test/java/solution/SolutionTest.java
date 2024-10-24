package solution;

import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.Arrays;
import java.util.stream.Stream;


public class SolutionTest {

    private static Stream<Arguments> tests() {
        return Stream.of(
                Arguments.of(new char[][]{
                    {'5','3','.','.','7','.','.','.','.'},
                    {'6','.','.','1','9','5','.','.','.'},
                    {'.','9','8','.','.','.','.','6','.'},
                    {'8','.','.','.','6','.','.','.','3'},
                    {'4','.','.','8','.','3','.','.','1'},
                    {'7','.','.','.','2','.','.','.','6'},
                    {'.','6','.','.','.','.','2','8','.'},
                    {'.','.','.','4','1','9','.','.','5'},
                    {'.','.','.','.','8','.','.','7','9'}
                }, true),

                // two 9's in first top-left 3x3 block
                Arguments.of(new char[][]{
                    {'5','3','.','.','7','.','.','.','.'},
                    {'6','.','9','1','9','5','.','.','.'},
                    {'.','9','.','.','.','.','.','6','.'},
                    {'8','.','.','.','6','.','.','.','3'},
                    {'4','.','.','8','.','3','.','.','1'},
                    {'7','.','.','.','2','.','.','.','6'},
                    {'.','.','.','.','.','.','2','8','.'},
                    {'6','.','.','4','1','9','.','.','5'},
                    {'.','.','.','.','8','.','.','7','9'}
                }, false),

                // two sixes in first column
                Arguments.of(new char[][]{
                    {'5','3','.','.','7','.','.','.','.'},
                    {'6','.','.','1','9','5','.','.','.'},
                    {'.','9','8','.','.','.','.','6','.'},
                    {'8','.','.','.','6','.','.','.','3'},
                    {'4','.','.','8','.','3','.','.','1'},
                    {'7','.','.','.','2','.','.','.','6'},
                    {'.','.','.','.','.','.','2','8','.'},
                    {'6','.','.','4','1','9','.','.','5'},
                    {'.','.','.','.','8','.','.','7','9'}
                }, false),

                Arguments.of(new char[][]{
                    {'.','.','.','.','5','.','.','1','.'},
                    {'.','4','.','3','.','.','.','.','.'},
                    {'.','.','.','.','.','3','.','.','1'},
                    {'8','.','.','.','.','.','.','2','.'},
                    {'.','.','2','.','7','.','.','.','.'},
                    {'.','1','5','.','.','.','.','.','.'},
                    {'.','.','.','.','.','2','.','.','.'},
                    {'.','2','.','9','.','.','.','.','.'},
                    {'.','.','4','.','.','.','.','.','.'}
                }, false)

        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testValidSudoku(char[][] board, boolean expected) {
        assertEquals(expected, Solution.validSudoku(board), () -> "Failed on board: " + Arrays.deepToString(board));
    }
}
