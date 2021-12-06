package easy

func twoSum(nums []int, target int) []int {
	tMap := make(map[int]int)
	for index, value := range nums {
		if v, ok := tMap[value]; !ok {
			tMap[value] = index
		} else {
			if target-value == value {
				return []int{index, v}
			}
		}
	}
	for k, v := range tMap {
		if i, ok := tMap[target-k]; ok && i != v {
			return []int{v, i}
		}
	}
	return nil
}
