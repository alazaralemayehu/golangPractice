package datastructure

import (
	"fmt"
	"sync"
)

type TreeNode struct {
	Key       int
	Value     int
	LeftNode  *TreeNode
	RightNode *TreeNode
}

type BinarySearchTree struct {
	RootNode *TreeNode
	Locker   sync.RWMutex
}

func (tree *BinarySearchTree) InsertElement(Key int, Value int) {
	tree.Locker.Lock()
	defer tree.Locker.Unlock()

	var treeNode *TreeNode
	treeNode = &TreeNode{Key, Value, nil, nil}
	if tree.RootNode == nil {

		tree.RootNode = treeNode

	} else {
		insertTreeNode(tree.RootNode, treeNode)
		// fmt.Println(tree.RootNode.LeftNode)
		// fmt.Println(tree.RootNode)
		// fmt.Println(tree.RootNode.RightNode)
	}
}

func insertTreeNode(RootNode *TreeNode, newTreeNode *TreeNode) {

	if RootNode.Key > newTreeNode.Key {
		if RootNode.LeftNode == nil {
			RootNode.LeftNode = newTreeNode
			// fmt.Println(RootNode.RightNode, RootNode.LeftNode)

		} else {
			insertTreeNode(RootNode.LeftNode, newTreeNode)
		}
	} else {
		if RootNode.RightNode == nil {
			RootNode.RightNode = newTreeNode
			// fmt.Println(RootNode.RightNode, RootNode.LeftNode)

		} else {
			insertTreeNode(RootNode.RightNode, newTreeNode)
		}
	}
}

func (tree *BinarySearchTree) preOrderTraverse(function func(int)) {
	tree.Locker.Lock()
	defer tree.Locker.Unlock()
	preOrderTraverseTree(tree.RootNode, function)
}

func preOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		function(treeNode.Value)
		preOrderTraverseTree(treeNode.LeftNode, function)
		preOrderTraverseTree(treeNode.RightNode, function)
	}
}

func (tree *BinarySearchTree) inOrderTraverse(function func(int)) {
	tree.Locker.Lock()
	defer tree.Locker.Unlock()

	inOrderTraverseTree(tree.RootNode, function)
}

func inOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		inOrderTraverseTree(treeNode.LeftNode, function)
		function(treeNode.Value)
		inOrderTraverseTree(treeNode.RightNode, function)
	}
}

func (tree *BinarySearchTree) postOrderTraverse(function func(int)) {
	tree.Locker.Lock()
	defer tree.Locker.Unlock()
	postOrderTraverseTree(tree.RootNode, function)
}

func postOrderTraverseTree(treeNode *TreeNode, function func(int)) {
	if treeNode != nil {
		postOrderTraverseTree(treeNode.LeftNode, function)
		postOrderTraverseTree(treeNode.RightNode, function)
		function(treeNode.Value)

	}
}

func (tree *BinarySearchTree) minNode() *int {
	defer tree.Locker.Lock()
	defer tree.Locker.Unlock()

	var treeNode *TreeNode
	treeNode = tree.RootNode

	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.LeftNode == nil {
			return &treeNode.Value
		}
		treeNode = treeNode.LeftNode
	}
}

func (tree *BinarySearchTree) maxNode() *int {
	defer tree.Locker.Lock()
	defer tree.Locker.Unlock()

	var treeNode *TreeNode
	treeNode = tree.RootNode

	if treeNode == nil {
		return (*int)(nil)
	}
	for {
		if treeNode.RightNode == nil {
			return &treeNode.Value
		}
		treeNode = treeNode.RightNode
	}
}

func (tree *BinarySearchTree) searchNode(Key int) bool {
	tree.Locker.Lock()
	defer tree.Locker.Unlock()

	return SearchNode(tree.RootNode, Key)
}

func SearchNode(treeNode *TreeNode, Key int) bool {
	if treeNode == nil {
		return false
	}

	if Key < treeNode.Key {
		return SearchNode(treeNode.LeftNode, Key)
	}
	if Key > treeNode.Key {
		return SearchNode(treeNode.RightNode, Key)
	}

	return true
}

//This code may not work, it will be fixed later
func (tree *BinarySearchTree) removeNode(Key int) {
	tree.Locker.Lock()
	defer tree.Locker.Unlock()

	removeNode(tree.RootNode, Key)
}

func removeNode(treeNode *TreeNode, Key int) {
	if Key < treeNode.Key {
		removeNode(treeNode.LeftNode, Key)
	}
	if Key > treeNode.Key {
		removeNode(treeNode.RightNode, Key)
	}

	if treeNode.LeftNode == nil && treeNode.RightNode == nil {
		treeNode = nil

		return
	}
	if treeNode.RightNode == nil {
		treeNode = treeNode.LeftNode
		return
	}
	if treeNode.LeftNode == nil {
		treeNode = treeNode.RightNode
		return
	}

	if treeNode.LeftNode != nil && treeNode.RightNode != nil {

		var node *TreeNode = treeNode.LeftNode
		for {

			if node.LeftNode != nil {
				node = node.LeftNode
			} else {
				break

			}
		}

		node.RightNode = treeNode.RightNode
		treeNode = node
		fmt.Print(treeNode.RightNode)
	}
}
