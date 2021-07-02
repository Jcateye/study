package main

import (
	"fmt"
)

func main() {
	nums := []int{7, 11, 15, 2, 5, 14, 3, 9, 8, 30, 20, 21, 19, 40}
	target := 19
	// 1
	fmt.Println(twoSum(nums, target))
	fmt.Println(twoSum2(nums, target))
	l1 := &ListNode{Val: 2}
	l2 := &ListNode{Val: 5}
	l1.addNode(4).addNode(3).addNode(9)
	l2.addNode(6).addNode(4).addNode(6)
	l1.PrintNode()
	l2.PrintNode()
	// 2
	addTwoNumbers(l1, l2).PrintNode()
	l3 := l1.reversal()
	l3.PrintNode()
	my := MergeSort(nums[:])
	fmt.Printf("result --> %v", my)
}

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
// 1. 两数相加1
func twoSum(nums []int, target int) []int {
	result := []int{0, 0}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				return result
			}

		}
	}
	return result
}

// 1. 两数相加2
func twoSum2(nums []int, target int) []int {
	result := []int{0, 0}
	maps := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		maps[nums[i]] = i
		a := target - nums[i]
		if _, ok := maps[a]; ok {
			result[0] = i
			result[1] = maps[a]
		}
	}
	return result
}

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

 

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func (node *ListNode) addNode(val int) (n *ListNode) {
	next := &ListNode{Val: val}
	cur := node
	for cur.Next != nil {
		cur = cur.Next
	}
	cur.Next = next
	return node
}

func (node *ListNode) PrintNode() {
	cur := &ListNode{Next: node}

	for cur.Next != nil {
		cur = cur.Next
		fmt.Printf("%v ", cur.Val)
	}
	fmt.Println()
}

// pre  -- cur -- next -- next.next
// 0 -- 1 -- 2 -- 3
//
func (node *ListNode) reversal() (n *ListNode) {
	var pre *ListNode
	cur := node
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
		pre.PrintNode()
	}
	return pre
}

/**
给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。

请你将两个数相加，并以相同形式返回一个表示和的链表。

你可以假设除了数字 0 之外，这两个数都不会以 0 开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	pre := ListNode{}
	cur := &pre
	carry := 0
	for l1 != nil || l2 != nil {
		x, y := 0, 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		sum := x + y + carry
		carry = sum / 10
		remainder := sum % 10
		next := ListNode{}
		next.Val = remainder
		cur.Next = &next
		cur = cur.Next
	}
	if carry > 0 {
		next := ListNode{Val: carry}
		cur.Next = &next
	}
	return pre.Next
}

// 归并排序 拆
func MergeSort(sourceArr []int) (temp []int) {
	mid := len(sourceArr) / 2
	left := MergeSort(sourceArr[:mid])
	right := MergeSort(sourceArr[mid:])
	return Merge(left, right)
}
// 归并排序 合
func Merge(left []int, right []int) (temp []int) {
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		} else {
			temp = append(temp, right[j])
			j++
		}
	}
	temp = append(temp, left[i:]...)
	temp = append(temp, right[j:]...)
	return
}
