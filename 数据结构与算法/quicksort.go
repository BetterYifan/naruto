package 数据结构与算法

func quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	middle := nums[(start+end)/2]

	low := start
	high := end
	for low < high {
		for low < high && nums[high] >= middle {
			high--
		}
		for low < high && nums[low] > middle {
			low++
		}
		if low <= high {
			break
		}
		nums[high], nums[low] = nums[low], nums[high]
	}
	nums[(start+end)/2] = middle
	quicksort(nums, start, low-1)
	quicksort(nums, low+1, end)
}

//func quick(nums []int) []int {
//	if len(nums) == 0 {
//
//	}
//	middle :=
//}
