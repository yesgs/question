package q0023

//合并K个升序链表

//给你一个链表数组，每个链表都已经按升序排列。

//请你将所有链表合并到一个升序链表中，返回合并后的链表。

//示例 1：
//
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	k := len(lists)
	dummyHead := &ListNode{}
	tail := dummyHead
	var minNode *ListNode

	var i = 0
	var minIndex = -1 //最小的链表下标

	for {

		//每次循环
		//最小下表不存在
		//最小节点不存在
		minIndex = -1
		minNode = nil

		//循环k个链表
		for i = 0; i < k; i++ {
			if lists[i] == nil {
				//如果某个链表为空，则退出本次循环
				continue
			}
			if minNode == nil {
				//第一次给最小的节点和下标赋值
				minNode = lists[i]
				minIndex = i
			}
			if lists[i].Val < minNode.Val {
				//比较当前链表和已知最小节点
				//获取最小的那个节点
				minNode = lists[i]
				minIndex = i
			}
		}
		//循环k个链表后未找到最小的
		if minIndex == -1 {
			//退出主循环
			break
		}
		//把最小的节点连起来
		tail.Next = minNode
		tail = tail.Next

		//最小的链表往前进一格
		lists[minIndex] = lists[minIndex].Next
	}

	return dummyHead.Next
}

func mergeKLists2(lists []*ListNode) *ListNode {
	var res *ListNode

	for i := 0; i < len(lists); i++ {
		res = mergeTwoLists(res, lists[i])
	}

	return res
}

func SliceToListNodeSlice(nodeList [][]int) []*ListNode {
	var ret []*ListNode

	for i := 0; i < len(nodeList); i++ {
		nodes := nodeList[i]

		var head = &ListNode{
			Val:  0,
			Next: nil,
		}
		var p = head

		for j := 0; j < len(nodes); j++ {
			p.Next = &ListNode{
				Val:  nodes[j],
				Next: nil,
			}
			p = p.Next
		}

		ret = append(ret, head.Next)
	}

	return ret
}

func ListNodeToSlice(list *ListNode) []int {
	var ret []int

	for list != nil {
		v := list.Val
		ret = append(ret, v)
		list = list.Next
	}

	return ret
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val <= list2.Val {
		//list1 比较小
		//list1.next 的下一个也挑一个小的
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	}
}
