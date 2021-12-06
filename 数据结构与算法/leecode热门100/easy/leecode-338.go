package easy

func countBits(n int) []int {
	res := make([]int, n+1)
	res[0] = 0
	for i := 1; i <= n; i++ {
		// 奇数
		if i&1 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}
