package solution;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;

public class SolutionTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
          Arguments.of("racecar", "carrace", true),
          Arguments.of("jar", "jam", false),
          Arguments.of("maj", "jam", true)
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testRomanToInt(String s, String t, boolean expected) {
        assertEquals(expected, Solution.isAnagram(s, t));
    }
}
