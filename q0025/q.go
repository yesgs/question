package q0025

import "question/common"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *common.ListNode, k int) *common.ListNode {
	node := head
	for i := 0; i < k; i++ {
		if node == nil {
			return head
		}
		node = node.Next
	}
	newHead := reverse(head, node)
	head.Next = reverseKGroup(node, k)
	return newHead
}

func reverse(first *common.ListNode, last *common.ListNode) *common.ListNode {
	prev := last
	for first != last {
		tmp := first.Next
		first.Next = prev
		prev = first
		first = tmp
	}
	return prev
}

func reverseKGroup2(head *common.ListNode, k int) *common.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//定义一个假的节点。
	dummy := &common.ListNode{}

	//假节点的next指向head。
	// dummy->1->2->3->4->5
	dummy.Next = head
	//初始化pre和end都指向dummy。
	//pre指每次要翻转的链表的头结点的上一个节点。
	var pre *common.ListNode = dummy
	//end指每次要翻转的链表的尾节点
	var end *common.ListNode = dummy
	for end.Next != nil {
		//循环k次，找到需要翻转的链表的结尾,这里每次循环要判断end是否等于空,因为如果为空，end.next会报空指针异常。
		//dummy->1->2->3->4->5 若k为2，循环2次，end指向2
		for i := 0; i < k && end != nil; i++ {
			end = end.Next
		}

		//如果end==null，即需要翻转的链表的节点数小于k，不执行翻转。
		if end == nil {
			break
		}

		//先记录下end.next,方便后面链接链表
		next := end.Next

		//然后断开链表
		end.Next = nil

		//记录下要翻转链表的头节点
		start := pre.Next

		//翻转链表,pre.next指向翻转后的链表。1->2 变成2->1。 dummy->2->1
		pre.Next = reverse2(start)

		//翻转后头节点变到最后。通过.next把断开的链表重新链接。
		start.Next = next

		//将pre换成下次要翻转的链表的头结点的上一个节点。即start
		pre = start
		
		//翻转结束，将end置为下次要翻转的链表的头结点的上一个节点。即start
		end = start
	}
	return dummy.Next
}

func reverse2(head *common.ListNode) *common.ListNode {
	//单链表为空或只有一个节点，直接返回原单链表
	if head == nil || head.Next == nil {
		return head
	}

	//前一个节点指针
	var preNode *common.ListNode
	//下一个节点指针
	var nextNode *common.ListNode
	//当前节点指针
	var curNode = head

	for curNode != nil {
		//nextNode 指向下一个节点,保存当前节点后面的链表。
		nextNode = curNode.Next

		//将当前节点指向前一个节点   null<-1<-2<-3<-4
		curNode.Next = preNode

		//preNode 指针向后移动。
		//preNode指向当前节点。
		preNode = curNode

		//curNode指针向后移动。
		//下一个节点变成当前节点
		curNode = nextNode
	}
	return preNode
}
