package 数据结构与算法

//给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
//
//具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。
//
// 
//
//示例 1：
//
//输入："abc"
//输出：3
//解释：三个回文子串: "a", "b", "c"
//示例 2：
//
//输入："aaa"
//输出：6
//解释：6个回文子串: "a", "a", "a", "aa", "aa", "aaa"

//动态规划:
//	状态：dp[i][j]表示s在[i,j]的区间是否是一个子串
//	状态转移方程：s[i]=s[j]&&(j-i<2||dp[i+1][j-1])

func countSubString(s string) int {
	if len(s) == 0 {
		return 0
	}
	var (
		size = len(s)
		dp   = make([][]bool, size)
		num  = 0
	)
	for j := 0; j < size; j++ {
		dp[j] = make([]bool, size)
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i < 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				num++
			}
		}
	}
	return num
}
