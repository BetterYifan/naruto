package middle

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int, 0)

	var dfs func(sum int, comb []int, idx int)

	dfs = func(sum int, comb []int, idx int) {
		if sum == target {
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			res = append(res, tmp)
			return
		}

		if sum > target {
			return
		}

		for i := idx; i < len(candidates); i++ {
			sum += candidates[i]
			comb = append(comb, candidates[i])
			idx = i
			dfs(sum, comb, idx)
			sum = sum - candidates[i]
			comb = comb[:len(comb)-1]
		}

	}

	dfs(0, []int{}, 0)
	return res
}
