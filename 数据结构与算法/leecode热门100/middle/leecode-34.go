package middle

func searchRange(nums []int, target int) []int {
	length := len(nums)
	left, right := 0, length-1
	res := make([]int, 0)

	for left < right {
		middle := (left + right) / 2
		if middle < target {
			left = middle + 1
		} else if middle > target {
			right = middle - 1
		} else {
			if len(res) == 0 {
				res = append(res, middle, middle)
			}
			left, right = middle, middle
			for left >= 0 && nums[left] == target {
				res[0] = left
				left--
			}
			for right <= length-1 && nums[right] == target {
				res[1] = left
				right++
			}
			return res
		}
	}
	return []int{-1, -1}
}
