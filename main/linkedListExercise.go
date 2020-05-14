package main

import "fmt"

type Node struct {
	property int
	nextNode *Node
}

type LinkedList struct {
	headNode *Node
}

func (linkedList *LinkedList) AddToHead(value int) {
	node := Node{}
	node.property = value
	if linkedList.headNode != nil {
		node.nextNode = linkedList.headNode
	}
	linkedList.headNode = &node

}
func createNode(value int) Node {
	return Node{property: value}
}

func (linkedList *LinkedList) AddAfter(nodeValue int, targetValue int) {
	node := createNode(nodeValue)
	for l := linkedList.headNode; l != nil; l = l.nextNode {
		if l.property == targetValue {
			node.nextNode = l.nextNode
			l.nextNode = &node
		}
	}
}

func (linkedList *LinkedList) AddToTail(value int) {
	node := Node{property: value}
	l := linkedList.headNode
	for {
		if l.nextNode == nil {
			l.nextNode = &node
			break
		}
		l = l.nextNode
	}
}
func (linkedlist *LinkedList) PrintLinkedList() {
	for l := linkedlist.headNode; l != nil; l = l.nextNode {
		fmt.Println(l.property)
	}
}

func (linkedList *LinkedList) NodeWithValue(property int) *Node {
	var nodeWith *Node
	for l := linkedList.headNode; l != nil; l = l.nextNode {
		if l.property == property {
			nodeWith = l
		}
	}
	return nodeWith
}
