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
//二分查找
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2){
		return findMedianSortedArrays(nums2, nums1)
	}
	k, low, high, mid1, mid2 := (len(nums1)+len(nums2)+1)>>1, 0, len(nums1), 0, 0
	for low < high{
		mid1 := low + (high - low)>>1
		mid2 := k - mid1
		if nums1[mid1] < nums2[mid2-1]{
			low = mid1 +1
		} else {
			high = mid1
		}
	}
	mid1, mid2 = high, k - high
	mL, mR := 0, 0
	if mid1 == 0{
		mL = nums2[mid2 - 1]
	}else if mid2 == 0{
		mL = nums1[mid1 -1]
	}else{
		mL = max(nums1[mid1-1], nums2[mid2-1])
	}
	if (len(nums1)+len(nums2))&1 == 1{
		return float64(mL)
	}
	if mid1 == len(nums1){
		mR = nums2[mid2]
	}else if mid2 == len(nums2){
		mR = nums1[mid1]
	}else{
		mR = min(nums2[mid2], nums1[mid1])
	}
	return float64(mL+mR)/2
}
func max(a,b int) int{
	if a>b{
		return a
	}
	return b
}
func min(a,b int) int{
	if a<b{
		return a
	}
	return b
}


// 合并、排序、查找中位数
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
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