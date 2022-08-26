package main

import "fmt"

func main() {
	//主要原因是 append 后的 nums 是新数组，它并不会影响 main 函数中的 nums。如果最后不是赋值给 nums，而是使用 copy
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}

func rotate(nums []int, k int) {
	k = k % len(nums)
	nums = append(nums[len(nums)-k:], nums[0:len(nums)-k]...)
	//copy(nums, append(nums[len(nums)-k:], nums[0:len(nums)-k]...))
}
