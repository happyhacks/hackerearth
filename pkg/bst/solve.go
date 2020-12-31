package bst

import "fmt"

type node struct {
	val         int
	right, left *node
}

func (root *node) insert(val int) *node {
	if root == nil {
		return &node{val: val}
	}
	if val > root.val {
		root.right = root.right.insert(val)
	} else {
		root.left = root.left.insert(val)
	}
	return root
}

func (root *node) search(val int) *node {
	if root == nil {
		return nil
	}
	if val > root.val {
		return root.right.search(val)
	} else if val < root.val {
		return root.left.search(val)
	}
	return root
}

func (root *node) inorder() {
	if root == nil {
		return
	}
	root.left.inorder()
	fmt.Println(root.val)
	root.right.inorder()
}

func (root *node) preorder() {
	if root == nil {
		return
	}
	fmt.Println(root.val)
	root.left.preorder()
	root.right.preorder()
}

func (root *node) postorder() {
	if root == nil {
		return
	}
	root.left.postorder()
	root.right.postorder()
	fmt.Println(root.val)
}

func Solve() {
	var n, val int
	fmt.Scan(&n)
	if n <= 0 {
		return
	}
	fmt.Scan(&val)
	root := &node{val, nil, nil}
	for i := 1; i < n; i++ {
		fmt.Scan(&val)
		root = root.insert(val)
	}
	fmt.Scan(&val)
	root.search(val).preorder()
}
