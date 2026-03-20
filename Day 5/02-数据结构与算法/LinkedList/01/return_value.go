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

// 插入节点（在尾部）
func insertAtTail[T any](head *ListNode[T], val T) *ListNode[T] {
	newNode := &ListNode[T]{Val: val}

	if head == nil {
		return newNode
	}

	// 找到尾节点
	current := head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
	return head
}

// 插入节点（在指定索引位置）
// index 从 0 开始，index=0 表示插入到头部
func insertAtIndex[T any](head *ListNode[T], index int, val T) *ListNode[T] {
	// 索引无效
	if index < 0 {
		return head
	}

	// 在头部插入
	if index == 0 {
		return insertAtHead(head, val)
	}

	// 找到插入位置的前一个节点
	current := head
	for i := 0; i < index-1 && current != nil; i++ {
		current = current.Next
	}

	// 索引超出范围
	if current == nil {
		return head
	}

	// 插入新节点
	newNode := &ListNode[T]{Val: val, Next: current.Next}
	current.Next = newNode

	return head
}

// 插入节点（在指定值之后）
func insertAfterValue[T comparable](head *ListNode[T], target T, val T) *ListNode[T] {
	if head == nil {
		return head
	}

	// 查找目标节点
	current := head
	for current != nil && current.Val != target {
		current = current.Next
	}

	// 未找到目标节点
	if current == nil {
		return head
	}

	// 在目标节点后插入
	newNode := &ListNode[T]{Val: val, Next: current.Next}
	current.Next = newNode

	return head
}

// 查找节点（按值查找）- 返回节点指针
func findNode[T comparable](head *ListNode[T], val T) *ListNode[T] {
	current := head
	for current != nil {
		if current.Val == val {
			return current
		}
		current = current.Next
	}
	return nil
}

// 查找节点（按索引查找）- 返回节点指针
func findByIndex[T any](head *ListNode[T], index int) *ListNode[T] {
	if index < 0 {
		return nil
	}
	current := head
	for i := 0; i < index && current != nil; i++ {
		current = current.Next
	}
	return current
}

// 查找节点（按值查找）- 返回索引位置
func findIndex[T comparable](head *ListNode[T], val T) int {
	current := head
	index := 0
	for current != nil {
		if current.Val == val {
			return index
		}
		current = current.Next
		index++
	}
	return -1 // 未找到返回 -1
}

// 查找节点（按条件查找）- 使用自定义函数
func findByCondition[T any](head *ListNode[T], condition func(T) bool) *ListNode[T] {
	current := head
	for current != nil {
		if condition(current.Val) {
			return current
		}
		current = current.Next
	}
	return nil
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

	// 查找演示
	fmt.Println("\n=== 查找操作演示 ===")

	// 1. 按值查找
	fmt.Println("\n1. 按值查找:")
	node := findNode(intHead, 4)
	if node != nil {
		fmt.Printf("找到节点，值: %v, 地址: %p\n", node.Val, node)
	} else {
		fmt.Println("未找到节点")
	}

	// 2. 按索引查找
	fmt.Println("\n2. 按索引查找:")
	node2 := findByIndex(intHead, 2)
	if node2 != nil {
		fmt.Printf("索引 2 的节点，值: %v\n", node2.Val)
	} else {
		fmt.Println("索引超出范围")
	}

	// 3. 查找索引位置
	fmt.Println("\n3. 查找值的索引位置:")
	index := findIndex(intHead, 5)
	if index != -1 {
		fmt.Printf("值 5 在索引位置: %d\n", index)
	} else {
		fmt.Println("未找到该值")
	}

	// 4. 按条件查找（查找大于 3 的第一个节点）
	fmt.Println("\n4. 按条件查找（查找大于 3 的第一个节点）:")
	node3 := findByCondition(intHead, func(val int) bool {
		return val > 3
	})
	if node3 != nil {
		fmt.Printf("找到符合条件的节点，值: %v\n", node3.Val)
	} else {
		fmt.Println("未找到符合条件的节点")
	}

	// 字符串链表查找
	fmt.Println("\n=== 字符串链表查找 ===")
	stringHead := createList([]string{"a", "b", "c", "d", "e"})

	strNode := findNode(stringHead, "c")
	if strNode != nil {
		fmt.Printf("找到字符串节点: %v\n", strNode.Val)
	}

	strIndex := findIndex(stringHead, "d")
	fmt.Printf("字符串 'd' 在索引位置: %d\n", strIndex)

	// 插入操作演示
	fmt.Println("\n=== 插入操作演示 ===")
	testHead := createList([]int{1, 2, 3, 4, 5})

	// 在尾部插入
	fmt.Println("\n1. 在尾部插入 6:")
	testHead = insertAtTail(testHead, 6)
	printList(testHead)

	// 在索引位置插入
	fmt.Println("\n2. 在索引 2 位置插入 99:")
	testHead = insertAtIndex(testHead, 2, 99)
	printList(testHead)

	// 在指定值之后插入
	fmt.Println("\n3. 在值 4 之后插入 88:")
	testHead = insertAfterValue(testHead, 4, 88)
	printList(testHead)

	// 在头部插入（索引 0）
	fmt.Println("\n4. 在索引 0 位置插入 0:")
	testHead = insertAtIndex(testHead, 0, 0)
	printList(testHead)
}
