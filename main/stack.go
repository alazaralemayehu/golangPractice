package main

type Element struct {
	property int
}

type Stack struct {
	elements     []*Element
	elementCount int
}

func (element *Element) NewElement(value int) {
	element.property = value
}

func (stack *Stack) New() {
	stack.elements = make([]*Element, 0)
}

func (stack *Stack) Push(element *Element) {
	stack.elements = append(stack.elements, element)
	stack.elementCount += 1
}

func (stack *Stack) Pop() *Element {
	if len(stack.elements) < 0 {
		return nil
	} else {
		var returnElement *Element = stack.elements[len(stack.elements)-1]
		stack.elementCount -= 1
		return returnElement
	}
}
