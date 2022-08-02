/*
You are given the head of a linked list. Delete the middle node, and return the head of the modified linked list.

The middle node of a linked list of size n is the ⌊n / 2⌋th node from the start using 0-based indexing, where ⌊x⌋ denotes the largest integer less than or equal to x.

For n = 1, 2, 3, 4, and 5, the middle nodes are 0, 1, 1, 2, and 2, respectively.

Input: head = [1,3,4,7,1,2,6]
Output: [1,3,4,1,2,6]
*/

package main

import (
	"bytes"
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) PrintNodes() {
	nextNode := ln
	buf := bytes.NewBufferString("[")
	for nextNode.Next != nil {
		buf.WriteString(fmt.Sprint(nextNode.Val))
		buf.WriteRune(' ')
		nextNode = nextNode.Next
	}
	buf.WriteString(fmt.Sprint(nextNode.Val))
	buf.WriteRune(']')
	fmt.Println(buf.String())
}

func main() {
	nodeCnt := 4
	pLN := &ListNode{0, nil}
	var nextNode *ListNode = pLN
	for i := 0; i < nodeCnt-1; i++ {
		if nextNode.Next == nil {
			nextNode.Next = &ListNode{i + 1, nil}
			nextNode = nextNode.Next
		}
	}
	pLN.PrintNodes()
	newListNode := deleteMiddle(pLN)
	newListNode.PrintNodes()
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	cur, slow, fast := head, head, head
	for fast.Next != nil {
		cur = slow
		slow = slow.Next
		fast = fast.Next
		if fast.Next != nil {
			fast = fast.Next
		}
	}
	cur.Next = slow.Next
	return head
}
