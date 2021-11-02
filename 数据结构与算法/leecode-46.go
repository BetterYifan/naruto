package 数据结构与算法

//给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
//
//示例 1：
//
//输入：nums = [1,2,3]
//输出：[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
//示例 2：
//
//输入：nums = [0,1]
//输出：[[0,1],[1,0]]
//示例 3：
//
//输入：nums = [1]
//输出：[[1]]

//func permute(nums []int) [][]int {
//	res := make([][]int, 0)
//	return
//}

func twoSum(nums []int, target int) []int {
	tMap := make(map[int]int)
	for index, value := range nums {
		if v, ok := tMap[value]; !ok {
			tMap[value] = index
		} else {
			if target-v == value {
				return []int{index, tMap[value]}
			}
		}
	}
	for k, v := range tMap {
		if i, ok := tMap[target-k]; ok {
			return []int{v, i}
		}
	}
	return nil
}
