package easy

func hammingDistance(x int, y int) int {
	diff := x ^ y
	count := 0
	for diff != 0 {
		diff = diff & (diff - 1)
		count++
	}
	return count
}
