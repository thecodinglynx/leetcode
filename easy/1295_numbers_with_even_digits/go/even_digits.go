package leetcode

import "fmt"

// while this solution is O(n) for time complexity,
// it does (unnecessarily) have a O(n) space complexity
// whereas it could be as efficient as O(1).
// However, the solution as shown here is straightforward
// and easily readable. If the input array was extremely large
// or space is limited, a better solution would be to not
// convert each number to a string but instead determine
// the digits by repeatedly dividing the number by 10 which
// while slightly less readable would be more space efficient
func findNumbers(nums []int) int {
	var res int
	for _, v := range nums {
		if len(fmt.Sprint(v))%2 == 0 {
			res++
		}
	}
	return res
}
