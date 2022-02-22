package middle

// 自己用枚举实现,将nums映射到二进制位
//func subsets(nums []int) [][]int {
//	if len(nums) == 0 {
//		return [][]int{}
//	}
//	res := make([][]int, 0)
//
//	// 1->2, 2->4, 3->8
//	for i := 0; i < (1 << len(nums)); i++ {
//		selc := make([]int, 0)
//		// 将i转为二进制，选取位数为1的值
//		for k, v := range nums {
//			if i>>k&1 == 1 {
//				selc = append(selc, v)
//			}
//		}
//		res = append(res, selc)
//	}
//	return res
//}

// 遍历
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	res := make([][]int, 0)
	res = append(res, []int{})
	for _, num := range nums {
		for _, arry := range res {
			tmp := make([]int, len(arry))
			// 记住要用copy
			copy(tmp, arry)
			tmp = append(tmp, num)
			res = append(res, tmp)
		}
	}
	return res
}
