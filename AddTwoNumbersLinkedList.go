/* * LeetCode Problem: Add Two Numbers
 * Difficulty: Medium
 * Description: Given two non-empty linked lists representing two non-negative integers, the digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and retur

type ListNode struct {
	Val int
	Next *ListNode
}
*/

package LeetCodeSolutions

import (
	"fmt"
	"regexp"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
		//, errors.New("list node cannot be nil")
	}
	remainder := 0
	var prevNode *ListNode
	var firstNode *ListNode
	l1Iter, l2Iter := 0, 0
	for l1 != nil || l2 != nil {
		sum := 0
		switch {
		case l1Iter == l2Iter:
			sum = remainder + l1.Val + l2.Val
		case l1Iter > l2Iter:
			sum = remainder + l1.Val
		case l1Iter < l2Iter:
			sum = remainder + l2.Val
		}
		remainder = sum / 10
		digit := sum % 10
		newNode := &ListNode{Val: digit}
		if prevNode == nil {
			prevNode = newNode
			firstNode = newNode
		} else {
			prevNode.Next = newNode
			prevNode = newNode
		}
		l1, l1Iter = getNextNode(l1, l1Iter)
		l2, l2Iter = getNextNode(l2, l2Iter)
		if (l1 == nil && l2 == nil) && remainder > 0 {
			newNode = &ListNode{Val: remainder}
			prevNode.Next = newNode
			prevNode = newNode
		}
	}
	return firstNode
}

func getNextNode(l *ListNode, iter int) (*ListNode, int) {
	if l == nil || l.Next == nil {
		return nil, iter
	}
	return l.Next, iter + 1
}

func NumberStringToListNode(numStr string) *ListNode {
	if len(numStr) == 0 {
		return nil
		//, errors.New("empty or nil string")
	}

	const pattern = "^\\d+$"
	if !regexp.MustCompile(pattern).MatchString(numStr) {
		return nil
		//, errors.New("invalid number string")
	}

	var firstNode *ListNode
	var prevNode *ListNode

	digits := []rune(numStr)
	for i := len(digits) - 1; i >= 0; i-- {
		newDig, err := strconv.Atoi(string(digits[i]))
		if err != nil {
			return nil
			//, errors.New("incorrect number")
		}
		newNode := &ListNode{Val: newDig}
		if firstNode == nil && prevNode == nil {
			firstNode = newNode
			prevNode = newNode
		} else {
			prevNode.Next = newNode
			prevNode = newNode
		}
	}
	return firstNode
	//, nil
}

func PrintNumber(ln *ListNode) {
	fmt.Print(NumericString(ln))
}

func NumericString(ln *ListNode) string {
	numS := ""
	iLn := ln
	for iLn != nil {
		numS = strconv.Itoa(iLn.Val) + numS
		if iLn.Next != nil {
			iLn = iLn.Next
		} else {
			iLn = nil
		}
	}
	return numS
}
