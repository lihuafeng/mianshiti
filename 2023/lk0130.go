package main

import (
	"fmt"
)
/**
力扣算法题
给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
 */
func main() {
	//s := "pwwwkew"
	s := "pwwdevdfkl"
	//s := "abcabcbb"
	l := lengthOfLongestSubstring(s)
	fmt.Println(l)
}

func lengthOfLongestSubstring(s string) int {
	var (
		left, right, length int
		window = make(map[rune]int)
		str = []rune(s)
	)
	// 当窗口的右边界不超过字符串大小
	for right < len(s) {
		// 将右边界的字符加入窗口
		r := str[right]
		right++
		window[r]++
		// 当窗口中出现重复的字符，移动窗口的左边界，直到不存在重复字符，进而继续移动窗口
		for window[r] > 1 {
			l := str[left]
			left++
			window[l]--
		}
		fmt.Printf("left=%v, right=%v, length=%v \n", left, right,length)
		// 满足不重复子串条件，计算/保存最长不重复子串
		if right-left > length {
			length = right - left
		}
	}
	return length
}


