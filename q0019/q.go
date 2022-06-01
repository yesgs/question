package q0019

//删除链表的倒数第 N 个结点
//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToListNode(nodes []int) *ListNode {
	var head = &ListNode{}
	var p = head

	for i := 0; i < len(nodes); i++ {
		node := nodes[i]
		tmp := ListNode{
			Val:  node,
			Next: nil,
		}
		p.Next = &tmp
		p = p.Next
	}

	return head.Next
}

func ListNodeToSlice(head *ListNode) []int {
	var ret []int

	for head != nil {
		ret = append(ret, head.Val)
		head = head.Next
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var p = head
	length := 0
	for p != nil {
		p = p.Next
		length++
	}

	//输入：head = [1,2,3,4,5], n = 2
	//输出：[1,2,3,5]

	delIndex := length - n

	if delIndex == 0 {
		return head.Next
	}

	i := 0
	p = head
	for p != nil {
		if i+1 == delIndex {
			if p.Next != nil {
				p.Next = p.Next.Next
			} else {
				p.Next = nil
			}
		}
		i++
		p = p.Next
	}

	return head
}
