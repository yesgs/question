package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(records []int) *ListNode {

	var head = &ListNode{
		Val:  0,
		Next: nil,
	}

	var p = head

	for i := 0; i < len(records); i++ {
		p.Next = &ListNode{
			Val:  records[i],
			Next: nil,
		}
		p = p.Next
	}

	return head.Next
}

func ListNodeToSlice(list *ListNode) []int {
	var ret []int

	for list != nil {
		ret = append(ret, list.Val)
		list = list.Next
	}

	return ret
}
