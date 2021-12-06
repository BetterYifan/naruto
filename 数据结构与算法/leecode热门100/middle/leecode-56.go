package middle

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	res := make([][]int, 0)
	x0, y0 := intervals[0][0], intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if y0 < intervals[i][0] {
			res = append(res, []int{x0, y0})
			x0 = intervals[i][0]
			y0 = intervals[i][1]
		} else {
			// 前一个y大于后一个x
			if y0 < intervals[i][1] {
				y0 = intervals[i][1]
			}
		}
	}
	return res
}
