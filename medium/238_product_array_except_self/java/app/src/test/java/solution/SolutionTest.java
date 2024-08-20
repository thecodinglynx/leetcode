package solution;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import static org.junit.jupiter.api.Assertions.assertEquals;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;


public class SolutionTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
            Arguments.of(new int[]{1,2,4,6}, new int[]{48,24,12,8})
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testSolution(int[] nums, int[] expected) {
        assertArrayEquals(expected, Solution.productExceptSelf(nums));
    }
}