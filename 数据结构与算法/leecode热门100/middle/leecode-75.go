package middle

func sortColors(nums []int) {
	if len(nums) == 0 {
		return
	}
	index, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[index] = 0
			index++
		} else if nums[i] == 1 {
			count++
		}
	}
	for ; index < len(nums); index++ {
		if count > 0 {
			nums[index] = 1
			count--
			continue
		}
		nums[index] = 2
	}
	return
}
