package test.java.solution;

import main.java.solution.Solution;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.Arrays;
import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class SolutionTest {

    private static Stream<Arguments> tests() {
        return Stream.of(
            Arguments.of(new int[]{100,4,200,1,3,2}, 4),
            Arguments.of(new int[]{0,3,7,2,5,8,4,6,0,1}, 9),
            Arguments.of(new int[]{1,3,5,7,9}, 1),
            Arguments.of(new int[]{2,20,4,10,3,4,5}, 4)
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testLongestConsecutiveNLogN(int[] arr, int expected) {
        int actual = Solution.longestConsecutiveN(arr);
        assertEquals(
            expected,
            actual,
            () -> "Failed for input: " + Arrays.toString(arr) +
                    ", expected: " + expected +
                    ", but got: " + actual
        );
    }
}
