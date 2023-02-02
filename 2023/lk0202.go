package main

import (
	"fmt"
	"sort"
)
/**
力扣算法题
寻找两个正序数组的中位数
 */
func main() {
	nums1 := []int{1,2}
	nums2 := []int{3,4}
	z := findMedianSortedArrays(nums1, nums2)
	fmt.Println(z)
}
// 合并、排序、查找中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var mid float64
	//切片合并
	newSli := append(nums1, nums2...)
	sort.Ints(newSli)

	length := len(newSli)
	if length % 2==0{
		z := length / 2 - 1
		mid = float64(newSli[z]+newSli[z+1])/2
	}else{
		z := ((length +1)/2)-1
		mid = float64(newSli[z])
	}
	return mid
}