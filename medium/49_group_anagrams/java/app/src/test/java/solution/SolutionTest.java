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
            Arguments.of(new String[] {"eat", "tea", "tan", "ate", "nat", "bat"}, List.of(List.of("eat", "tea", "ate"), List.of("tan", "nat"), List.of("bat"))),
            Arguments.of(new String[] {""}, List.of(List.of("")))
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testGroupAnagrams(String[] strs, List<List<String>> expected) {
        List<List<String>> result = Solution.groupAnagrams(strs);
        Set<Set<String>> expectedSet = expected.stream()
            .map(Set::copyOf)
            .collect(Collectors.toSet());
        Set<Set<String>> resultSet = result.stream()
            .map(Set::copyOf)
            .collect(Collectors.toSet());

        assertEquals(expectedSet, resultSet);
    }
}
