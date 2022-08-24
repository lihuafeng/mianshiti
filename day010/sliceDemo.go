package main

import "fmt"

func main()  {
	//148
	//var nums1 []interface{}
	//nums2 := []int{1, 3, 4}
	//nums3 := append(nums1, nums2...)
	//fmt.Println(len(nums3))

	//147
	var nums1 []interface{}
	nums2 := []int{1, 3, 4}
	nums3 := append(nums1, nums2)
	fmt.Println(len(nums3))
	fmt.Println(nums3)
}
