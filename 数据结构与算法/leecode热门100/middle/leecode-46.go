package middle

import "sort"

func permute(nums []int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)

	var backtrace func(path []int, start int)

	backtrace = func(path []int, start int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}

		for i := start; i < len(nums); i++ {
			flag := false
			for _, v := range path {
				if v == nums[i] {
					flag = true
				}
			}
			if flag {
				continue
			}

			path = append(path, nums[i])
			backtrace(path, start)
			path = path[:len(path)-1]
		}

	}

	backtrace([]int{}, 0)

	return ans
}
