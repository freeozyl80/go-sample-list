package main

import (
	"fmt"
)

// 这里有个tricks, 就是将map[int][int] 类似 Set 去存 原始数据。不过缺陷是无论如何第一次都不能得到结果的
func twoSum(nums [4]int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if _, ok := m[another]; ok {
			return []int{m[another], i}
		}
		m[nums[i]] = i
	}
	return nil
}

func main() {
	var nums = [4]int{2, 7, 11, 15}
	var target = 9
	fmt.Println(twoSum(nums, target))
}

/*
 * 题目
 * Given nums = [2, 7, 11, 15], target = 9,
 *
 * Because nums[0] + nums[1] = 2 + 7 = 9,
 * return [0, 1]
 */
