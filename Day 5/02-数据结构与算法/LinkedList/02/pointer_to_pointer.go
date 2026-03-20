package main

import "fmt"

type ListNode[T any] struct {
	Val  T
	Next *ListNode[T]
}

func createList[T any](head **ListNode[T], values []T) {
	if len(values) == 0 {
		*head = nil
		return
	}
	*head = &ListNode[T]{Val: values[0]}
	current := *head
	fmt.Printf("首节点地址: %p, 值: %v\n", current, current.Val)
	for i := 1; i < len(values); i++ {
		current.Next = &ListNode[T]{Val: values[i]}
		current = current.Next
	}
	fmt.Printf("尾节点地址: %p, 值: %v\n", current, current.Val)
	printList(*head)
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

func insertAtHead[T any](head **ListNode[T], val T) {
	newNode := &ListNode[T]{Val: val, Next: *head}
	*head = newNode
}

func deleteNode[T comparable](head **ListNode[T], val T) {
	if *head == nil {
		return
	}
	if (*head).Val == val {
		*head = (*head).Next
		fmt.Println("\n删除节点:", val)
		printList(*head)
		return
	}
	current := *head
	for current.Next != nil && current.Next.Val != val {
		current = current.Next
	}
	if current.Next != nil {
		current.Next = current.Next.Next
	}
	fmt.Println("\n删除节点:", val)
	printList(*head)
}

func main() {
	// 整数链表
	fmt.Println("=== 指针的指针方式 ===")
	fmt.Println("\n整数链表:")
	var intHead *ListNode[int]
	createList(&intHead, []int{1, 2, 3, 4, 5})
	deleteNode(&intHead, 3)
	insertAtHead(&intHead, 0)
	fmt.Println("\n插入头节点 0 后:")
	printList(intHead)

	// 字符串链表
	fmt.Println("\n字符串链表:")
	var stringHead *ListNode[string]
	createList(&stringHead, []string{"a", "b", "c", "d", "e"})
	deleteNode(&stringHead, "a")
	insertAtHead(&stringHead, "z")
	fmt.Println("\n插入头节点 z 后:")
	printList(stringHead)
}
