package middle

var digMap = map[uint8]string{
	2: "abc",
	3: "def",
	4: "ghi",
	5: "jkl",
	6: "mno",
	7: "pqrs",
	8: "tuv",
	9: "wxyz",
}

var combines []string

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	backtrack(digits, 0, "")
	return combines
}

func backtrack(digits string, index int, combine string) {
	// 说明已经遍历到最后一位
	if index == len(digits) {
		combines = append(combines, combine)
		return
	}
	letter := digMap[digits[index]-'0']
	letterLen := len(letter)
	for i := 0; i < letterLen; i++ {
		backtrack(digits, index+1, combine+string(letter[i]))
	}
}
