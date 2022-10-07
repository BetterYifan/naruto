package middle

//G(n) = G(0)*G(n-1) + G(1) * (n-2)+...+G(n-1)*G(0), G(2) = G(0)*G(1)+G(1)*G(0)
func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
