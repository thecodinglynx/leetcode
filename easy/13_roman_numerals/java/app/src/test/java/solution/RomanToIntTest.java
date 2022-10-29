package solution;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;

public class RomanToIntTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
          Arguments.of("III", 3),
          Arguments.of("LVIII", 58),
          Arguments.of("MCMXCIV", 1994)
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testRomanToInt(String input, int expected) {
        assertEquals(expected, RomanToInt.romanToInt(input));
    }
}
