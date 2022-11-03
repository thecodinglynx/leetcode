package solution;

import org.junit.jupiter.params.ParameterizedTest;
import org.junit.jupiter.params.provider.Arguments;
import org.junit.jupiter.params.provider.MethodSource;

import static org.junit.jupiter.api.Assertions.assertEquals;

import java.util.stream.Stream;

public class ContainerWithMostWaterTest {
    
    private static Stream<Arguments> tests() {
        return Stream.of(
          Arguments.of(new int[]{1,1}, 1),
          Arguments.of(new int[]{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49),
          Arguments.of(new int[]{2, 3, 4, 5, 18, 17, 6}, 17),
          Arguments.of(new int[]{1, 1000, 1000, 6, 2, 5, 4, 8, 3, 7}, 1000)
        );
    }

    @ParameterizedTest
    @MethodSource("tests")
    void testContainerWithMostWater(int[] input, int expected) {
        assertEquals(expected, new ContainerWithMostWater().maxArea(input));
    }
}
