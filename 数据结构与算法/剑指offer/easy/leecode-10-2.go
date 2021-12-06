package easy

func numWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	const mod = 1e9 + 7
	// i:n-2, j:n-1
	i, j := 1, 1
	for k := 0; k < n; k++ {
		i, j = j, (i+j)%mod
	}
	return i
}
