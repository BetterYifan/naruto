package easy

func findRepeatNumber(nums []int) int {
	tmp := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := tmp[num]; ok {
			return num
		}
		tmp[num] = struct{}{}
	}
	return 0
}
