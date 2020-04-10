package linklist_demo

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

type DListNode struct {
	Value interface{}
	Pre   *DListNode
	Next  *DListNode
}

//创建单向链表
func CreateNode(node *ListNode, max int)  {
	cur := node
	for i:=1; i< max; i++ {
		cur.Next = &ListNode{}
		cur.Next.Value = i
		cur = cur.Next
	}
}

//单向链表翻转
func Reverse(head *ListNode) *ListNode {
	var pre *ListNode //定义一个节点  值为nil
	pre = nil
	for head != nil{
		next := head.Next //保存头结点的下一个节点
		head.Next = pre   //将当前的next节点更新为pre节点
		pre = head		  //更新pre节点为当前节点
		head = next		  //更新头结点为下一个节点
	}
	return pre
}

//双向链表翻转
func DoubleReverse(list *DListNode) *DListNode  {
	var pre, next *DListNode //定义一个节点  值为nil
	pre, next = nil, nil
	for list != nil{
		next = list.Next
		list.Pre = next
		list.Next = pre
		pre = list
		list = next
	}
	return pre
}

//打印链表
func PrintNode(node *ListNode)  {
	for node != nil {
		fmt.Print(node.Value, "->")
		node = node.Next
	}
	fmt.Println()
}
//打印链表
func PrintDNode(node *DListNode)  {
	for node != nil {
		fmt.Println(node, "\t")
		node = node.Next
	}
	fmt.Println()
}