func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, num := range nums {
		if value, found := m[target-num]; found {
			return []int{value, index}
		}

		m[num] = index
	}
	return nil
}
