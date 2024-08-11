package solution;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.Arrays;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;
import java.util.stream.Stream;


public class SolutionTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
            Arguments.of(new int[]{1,1,1,2,2,3}, 2, new int[]{1,2})
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testTopKFrequent(int[] nums, int k, int[] expected) {
        int[] result = Solution.topKFrequent(nums, k);
        Arrays.sort(result);
        Arrays.sort(expected);
        assertArrayEquals(expected, result);
    }
}
