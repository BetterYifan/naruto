package main

import (
	"fmt"
	"github.com/naruto/数据结构与算法/leecode热门100/middle"
)

func main() {
	a := [][]int{
		{1, 3, 1}, {1, 5, 1}, {4, 2, 1},
	}
	fmt.Println(middle.MinPathSum(a))
}
