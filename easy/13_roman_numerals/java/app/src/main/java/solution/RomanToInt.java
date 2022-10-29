package solution;

import java.util.Map;

public class RomanToInt {
    public static int romanToInt(String s) {
        final Map<String, Integer> m = Map.of(
            "I", 1,
            "V", 5,
            "X", 10,
            "L", 50,
            "C", 100,
            "D", 500,
            "M", 1000
        );
        
        Integer nr = 0;
        Integer prev = 0;
        for (int i = 0; i < s.length(); i++) {
            var curVal = m.get("" + s.charAt(i));
            switch(prev) {
                case 1:
                    if (curVal == 5 || curVal == 10) {
                        nr += (curVal - 1 - prev);
                    } else {
                        nr += curVal;
                    }
                    break;
                case 10:
                    if (curVal == 50 || curVal == 100) {
                        nr += (curVal - 10 - prev);
                    } else {
                        nr += curVal;
                    }
                    break;
                case 100:
                    if (curVal == 500 || curVal == 1000) {
                        nr += (curVal - 100 - prev);
                    } else {
                        nr += curVal;
                    }
                    break;
                default:
                    nr += curVal;
                    break;
            }
            
            prev = curVal;
        }
        
        return nr;
    }
}