package dom

import "dingo/components"

type DOMNode struct {
	data     components.Component
	children []*DOMNode
}

type DOMTree struct {
	root *DOMNode
}

func (node *DOMNode) AddChild(cmp components.Component) {
	newNode := &DOMNode{data: cmp}
	node.children = append(node.children, newNode)
}

func Render() {

}
