package middle

// 双指针代表起是位置
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	loc := make(map[uint8]int)
	l, r, max := 0, 1, 1
	loc[s[l]] = l
	for l < r && r < len(s) {
		if v, ok := loc[s[r]]; ok {
			l = v + 1
		}
		if r-l+1 > max {
			max = r - l + 1
		}
		loc[s[r]] = r
		r++
	}
	return max
}
