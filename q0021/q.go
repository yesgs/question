package q0021

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//非递归
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var head = &ListNode{
		Val:  0,
		Next: nil,
	}

	var p = head
	for list1 != nil && list2 != nil {
		//p 指针在两条链上反复横跳
		//谁小谁前进一个
		if list1.Val <= list2.Val {
			p.Next = list1
			p = p.Next
			list1 = list1.Next
		} else {
			p.Next = list2
			p = p.Next
			list2 = list2.Next
		}
	}
	if list1 != nil {
		p.Next = list1
	}
	if list2 != nil {
		p.Next = list2
	}
	return head.Next
}

//递归
func mergeTwoLists2(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		//list1 比较小
		//list1.next 的下一个也挑一个小的
		list1.Next = mergeTwoLists2(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists2(list1, list2.Next)
		return list2
	}
}

func q0021(list1 *ListNode, list2 *ListNode) *ListNode {
	//return mergeTwoLists(list1, list2)
	return mergeTwoLists2(list1, list2)
}

func SliceToListNode(records []int) (list *ListNode) {
	var ret *ListNode = &ListNode{
		Val:  0,
		Next: nil,
	}
	var p *ListNode = ret

	for i := 0; i < len(records); i++ {
		record := records[i]

		p.Next = &ListNode{
			Val:  record,
			Next: nil,
		}
		p = p.Next
	}

	return ret.Next
}

func ListNodeToSlice(list *ListNode) []int {
	var ret []int

	for list != nil {
		ret = append(ret, list.Val)
		list = list.Next
	}

	return ret
}
