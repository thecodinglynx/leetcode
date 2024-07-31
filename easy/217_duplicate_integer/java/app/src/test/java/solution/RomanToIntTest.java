package solution;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;

public class RomanToIntTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
          Arguments.of(new int[]{1,2,3,3}, true),
          Arguments.of(new int[]{1,2,3,4}, false)
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testRomanToInt(int[] nums, boolean expected) {
        assertEquals(expected, Solution.hasDuplicate(nums));
    }
}
