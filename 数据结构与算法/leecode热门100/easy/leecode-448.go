package easy

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	lack := make([]int, 0)
	tmp := make(map[int]struct{})

	for _, i := range nums {
		tmp[i] = struct{}{}
	}
	for i := 1; i <= n; i++ {
		if _, ok := tmp[i]; !ok {
			lack = append(lack, i)
		}
	}
	return lack
}
