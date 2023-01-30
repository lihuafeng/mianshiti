package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}
/**
两数相加 进位计算
 */
func main() {
	l1 := &ListNode{2,&ListNode{4, &ListNode{3,&ListNode{}}}}
	l2 := &ListNode{5,&ListNode{6, &ListNode{4,&ListNode{}}}}
	l :=addTwoNumbers(l1,l2)
	fmt.Printf("%#v", l)
}
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}

	var l1Val, l2Val int
	if l1 != nil{
		l1Val = l1.Val
		l1 = l1.Next
	}
	if l2 != nil{
		l2Val = l2.Val
		l2 = l2.Next
	}
	res.Val = l1Val + l2Val
	if res.Val > 9{
		res.Val = res.Val % 10
		if l1==nil{
			l1.Val++
		}else if l2==nil{
			l2.Val++
		}else{
			res.Next = &ListNode{
				Val:1,
			}
		}
	}
	if l1 ==nil && l2 ==nil{
		return res
	}
	res.Next = addTwoNumbers(l1, l2)
	return res
}
