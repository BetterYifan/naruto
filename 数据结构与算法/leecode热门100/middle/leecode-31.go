package middle

// 1. 找到较小的左边和较小的右边
func nextPermutation(nums []int) {
	n := len(nums)
	left := n - 2
	for left >= 0 && nums[left] >= nums[left+1] {
		left--
	}
	if left >= 0 {
		right := n - 1
		for right >= 0 && nums[left] >= nums[right] {
			right--
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
}

func reverse(num []int) {

}
