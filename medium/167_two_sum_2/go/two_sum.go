package leetcode

// problem asked to use constant space. Best solution for this
// therefore is to use two pointers.
// Otherwise, another possible solution would be to use a hashmap.
func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 2 {
		return []int{1, 2}
	}
	p1 := 0
	p2 := p1 + 1
	for {
		diff := target - numbers[p1]
		if diff == numbers[p2] {
			return []int{p1 + 1, p2 + 1}
		}
		if diff > numbers[p2] {
			if p2 == len(numbers)-1 {
				p1 = p1 + 1
				p2 = p1 + 1
			} else {
				p2 = p2 + 1
			}
		} else {
			p1 = p1 + 1
			p2 = p1 + 1
		}
	}
}
