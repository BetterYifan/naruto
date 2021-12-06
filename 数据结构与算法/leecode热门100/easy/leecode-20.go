package easy

var (
	hash = map[byte]byte{
		']': '[',
		'}': '{',
		')': '(',
	}
)

func isValid(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if hash[s[i]] > 0 {
			// 前一个不等于
			if len(stack) == 0 || stack[len(stack)-1] != hash[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			// 左括号，放入站
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}
