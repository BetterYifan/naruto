package middle

func generateParenthesis(n int) []string {
	res := make([]string, 0)
	if n == 0 {
		return res
	}

	if n == 1 {
		return []string{"()"}
	}

	resMap := combine(n, 2, map[string]struct{}{"()": {}})
	for k := range resMap {
		res = append(res, k)
	}
	return res
}

func combine(n int, count int, cur map[string]struct{}) map[string]struct{} {
	if count > n {
		return cur
	}
	res := make(map[string]struct{})

	for c := range cur {
		for i := 1; i <= len(c); i++ {
			res[c[0:i]+"()"+c[i:]] = struct{}{}
		}
	}
	return combine(n, count+1, res)
}
