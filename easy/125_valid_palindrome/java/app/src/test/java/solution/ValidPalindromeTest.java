package test.java.solution;

import main.java.solution.ValidPalindrome;
import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import java.util.stream.Stream;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class ValidPalindromeTest {

    private static Stream<Arguments> tests() {
        return Stream.of(
            Arguments.of("Was it a car or a cat I saw?", true),
            Arguments.of("tab a cat?", false),
            Arguments.of("", true),
            Arguments.of("ab,,,,,b.@%@#*!@$#a?", true)
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testValidSudoku(final String str, boolean expected) {
        assertEquals(expected, ValidPalindrome.validate(str), () -> "Failed on string: " + str);
    }
}
