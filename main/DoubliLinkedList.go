package main

import "fmt"

type DoubliLinkedNode struct {
	property     int
	nextNode     *DoubliLinkedNode
	previousNode *DoubliLinkedNode
}

type DoubliLinkedList struct {
	headNode *DoubliLinkedNode
}

func (d *DoubliLinkedList) AddToHead(value int) {
	dln := DoubliLinkedNode{property: value}
	if d.headNode == nil {
		d.headNode = &dln
	} else {
		d.headNode.previousNode = &dln
		dln.nextNode = d.headNode
		d.headNode = &dln
	}
}

func (d DoubliLinkedList) NodeBetweenValues(leftProperty, rightProperty int) *DoubliLinkedNode {
	for l := d.headNode; l != nil; l = l.nextNode {
		if l.nextNode != nil && l.previousNode != nil {
			if l.previousNode.property == leftProperty && l.nextNode.property == rightProperty {
				return l
			}
		}
	}
	return nil
}

func (d *DoubliLinkedList) AddAfter(nodeProperty int, value int) {
	for l := d.headNode; l != nil; l = l.nextNode {
		if l.property == nodeProperty {
			node := DoubliLinkedNode{property: value}
			node.nextNode = l.nextNode
			node.previousNode = l
			l.nextNode = &node
			break
		}
	}
}

func (LinkedList *DoubliLinkedList) PrintAll() {
	for l := LinkedList.headNode; l != nil; l = l.nextNode {
		fmt.Println(l.property)
	}
}
