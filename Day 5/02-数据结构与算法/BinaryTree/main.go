package main

import "fmt"

// 定义二叉树节点
type TreeNode[T any] struct {
	Val   T
	Left  *TreeNode[T]
	Right *TreeNode[T]
}

// 插入节点到二叉搜索树
func insertBST(root *TreeNode[int], val int) *TreeNode[int] {
	if root == nil {
		return newNode[int](val)
	}
	if val < root.Val {
		root.Left = insertBST(root.Left, val)
	} else if val > root.Val {
		root.Right = insertBST(root.Right, val)
	}
	return root
}

func searchBST(root *TreeNode[int], val int) *TreeNode[int] {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return searchBST(root.Left, val)
	} else if val > root.Val {
		return searchBST(root.Right, val)
	}
	return nil
}

// 创建节点
func newNode[T any](val T) *TreeNode[T] {
	return &TreeNode[T]{Val: val}
}

// 前序遍历（根-左-右）
func preorder[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}

// 中序遍历（左-根-右）
func inorder[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}
	inorder(root.Left)
	fmt.Print(root.Val, " ")
	inorder(root.Right)
}

// 后序遍历（左-右-根）
func postorder[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}
	postorder(root.Left)
	postorder(root.Right)
	fmt.Print(root.Val, " ")
}

// 层序遍历(广度优先)
func levelPrder[T any](root *TreeNode[T]) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	levelPrder(root.Left)
	levelPrder(root.Right)
}

func main() {
	root := newNode[int](1)
	root.Left = newNode[int](2)
	root.Right = newNode[int](3)
	root.Left.Left = newNode[int](4)
	root.Left.Right = newNode[int](5)
	root.Right.Left = newNode[int](6)
	root.Right.Right = newNode[int](7)

	// 前序遍历
	fmt.Print("前序遍历（根-左-右）: ")
	preorder(root)
	fmt.Println()

	// 中序遍历
	fmt.Print("中序遍历（左-根-右）: ")
	inorder(root)
	fmt.Println()

	// 后序遍历
	fmt.Print("后序遍历（左-右-根）: ")
	postorder(root)
	fmt.Println()

	// 二叉搜索树示例
	fmt.Println("\n二叉搜索树:")
	bst := newNode[int](5)
	bst = insertBST(bst, 3)
	bst = insertBST(bst, 7)
	bst = insertBST(bst, 1)
	bst = insertBST(bst, 9)
	bst = insertBST(bst, 4)
	bst = insertBST(bst, 6)

	fmt.Print("BST 中序遍历（应为升序）: ")
	inorder(bst)
	fmt.Println()

	// 层序遍历
	fmt.Print("BST 层序遍历（广度优先）: ")
	levelPrder(bst)
	fmt.Println()

	// 搜索节点
	fmt.Println("搜索节点 5: ", searchBST(bst, 5))
	fmt.Println("搜索节点 8: ", searchBST(bst, 8))
}
