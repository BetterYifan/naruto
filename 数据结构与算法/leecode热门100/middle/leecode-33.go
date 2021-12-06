package middle

// 二分查找
func search(nums []int, target int) int {
	length := len(nums)
	left, right := 0, length-1

	for left <= right {
		middle := (left + right) / 2
		if target == nums[middle] {
			return middle
		} else if nums[middle] < nums[right] {
			// 此时右边有序
			if nums[middle] < target && target <= nums[right] {
				left = middle + 1
			} else {
				right = middle - 1
			}
		} else {
			// 左边有序
			if target >= nums[left] && target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}
