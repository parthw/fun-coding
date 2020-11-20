package main

import "fmt"

// TreeNode for BST
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

func (t *TreeNode) insert(value int) {

	if value < t.val {
		if t.left == nil {
			t.left = &TreeNode{val: value}
			return
		}
		t.left.insert(value)
		return
	}

	if value > t.val {
		if t.right == nil {
			t.right = &TreeNode{val: value}
			return
		}
		t.right.insert(value)
		return
	}

	if t.val == value {
		return
	}
}

func (t *TreeNode) traverse() {
	if t.left != nil {
		t.left.traverse()
	}
	fmt.Println(t.val)
	if t.right != nil {
		t.right.traverse()
	}
}

func main() {
	values := []int{1, 2, 10, -9, -8, 3, 4, 4, 5, 6, 5, 6, 7, 8, 9, 0}

	tnode := &TreeNode{}
	// Inserting in BST
	for _, value := range values {
		tnode.insert(value)
	}

	// traversing Tree
	tnode.traverse()

}
