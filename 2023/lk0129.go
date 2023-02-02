package main

import "fmt"

/**
力扣算法题
 */
func main() {
	var nums = []int{2,5,11,5}
	var target = 10
	index := twoSum(nums, target)
	fmt.Print(index)
}
/**
执行用时：
32 ms
, 在所有 Go 提交中击败了
15.07%
的用户
内存消耗：
3.4 MB
, 在所有 Go 提交中击败了
87.62%
的用户
通过测试用例：
57 / 57
 */
func twoSum1(nums []int, target int) []int {
	index := len(nums)
	for i,_:=range nums {
		for j:=i;j<index-1;j++{
			if nums[i]+nums[j+1] == target{
				return []int{i, j+1}
			}
		}
	}
	return []int{}
}

func twoSum2(nums []int, target int) []int {
	index := len(nums)
	var scl int
	for i, _ := range nums{
		scl = target - nums[i]
		for j:=i;j<index-1;j++{
			if scl == nums[j+1]{
				return []int{i, j+1}
			}
		}
	}
	return []int{}
}
/**
利用map，把差值以key,value的关系存储
你可以想出一个时间复杂度小于 O(n2) 的算法吗？
 */
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		if preIndex, ok := m[target-val]; ok {
			return []int{preIndex, index}
		} else {
			m[val] = index
		}
	}
	return []int{}
}
