package solution;

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
            Arguments.of(List.of("neet", "code", "love", "you")),
            Arguments.of(List.of()),
            Arguments.of(List.of("whatisthishere", "3$$k4jk3#", "okwhataboutthis!@#$"))
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void encodeDecode(List<String> strs) {
        assertEquals(strs, Solution.encodeDecode(strs));
    }
}
