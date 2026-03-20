package main

import "fmt"

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func createList[T any](values []T) *ListNode[T] {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode[T]{Val: values[0]}
	current := head
	fmt.Printf("首节点地址: %p, 值: %v\n", current, current.Val)
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode[T]{Val: values[i]}
		current = current.Next
	}
	fmt.Printf("尾节点地址: %p, 值: %v\n", current, current.Val)
	printList(head)
	return head
}

func printList[T any](head *ListNode[T]) {
	current := head
	fmt.Print("链表内容:")
	for current != nil {
		fmt.Printf("%v -> ", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

// 插入节点（在头部）
func insertAtHead[T any](head *ListNode[T], val T) *ListNode[T] {
	newNode := &ListNode[T]{Val: val, Next: head}
	return newNode
}

func deleteNode[T comparable](head *ListNode[T], val T) *ListNode[T] {
	if head == nil {
		return nil
	}
	if head.Val == val {
		return head.Next
	}
	current := head
	for current.Next != nil && current.Next.Val != val {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	fmt.Println("\n删除节点:", val)
	printList(head)
	return head
}

func main() {
	// 整数链表
	fmt.Println("=== 返回值方式 ===")
	fmt.Println("\n整数链表:")
	intHead := createList([]int{1, 2, 3, 4, 5})
	intHead = deleteNode(intHead, 3)
	intHead = insertAtHead(intHead, 0)
	fmt.Println("\n插入头节点 0 后:")
	printList(intHead)

	// 字符串链表
	fmt.Println("\n字符串链表:")
	stringHead := createList([]string{"a", "b", "c", "d", "e"})
	stringHead = deleteNode(stringHead, "a")
	stringHead = insertAtHead(stringHead, "z")
	fmt.Println("\n插入头节点 z 后:")
	printList(stringHead)
}
