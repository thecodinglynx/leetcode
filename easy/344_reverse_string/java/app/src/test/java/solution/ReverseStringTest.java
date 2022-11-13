package solution;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertArrayEquals;

import java.util.stream.Stream;

public class ReverseStringTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
          Arguments.of(new char[]{'h', 'e', 'l', 'l', 'o'}, new char[]{'o', 'l', 'l', 'e', 'h'})
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testRomanToInt(char[] input, char[] expected) {
        final var rs = new ReverseString();
        rs.reverseString(input);
        System.out.println(input);
        assertArrayEquals(expected, input);
    }
}
