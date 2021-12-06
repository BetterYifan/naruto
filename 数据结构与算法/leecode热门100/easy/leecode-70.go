package easy

// 递归
func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	arr := make([]int, n+1, n+1)
	arr[0], arr[1] = 1, 1
	for i := 2; i < n; i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	return arr[n]
}
