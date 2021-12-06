package middle

func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	if len(height) == 1 {
		return height[0]
	}
	max := 0
	left, right := 1, len(height)
	for left < right {
		if max < (right-left)*min(height[left-1], height[right-1]) {
			max = (right - left) * min(height[left-1], height[right-1])

		}
		if height[left-1] < height[right-1] {
			left++
		} else {
			right--
		}

	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
