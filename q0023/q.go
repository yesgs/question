package q0023

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(records []int) *ListNode {
	var head = &ListNode{}
	var p = head

	for i := 0; i < len(records); i++ {
		record := records[i]
		p.Next = &ListNode{
			Val:  record,
			Next: nil,
		}
		p = p.Next
	}
	return head.Next
}

func ListNodeToSlice(node *ListNode) []int {
	if node == nil {
		return []int{}
	}
	var ret = []int{}
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}

	return ret
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	p1, p2 := head, head.Next

	for p1 != nil && p2 != nil {
		p1.Val, p2.Val = p2.Val, p1.Val
		p1 = p2.Next
		if p1 != nil {
			p2 = p1.Next
		} else {
			p2 = nil
		}
	}
	return head
}
