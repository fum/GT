package Components

import "fmt"

func NewNode(name string) *Node {
	return &Node{Transform: NewTransform(), Name: name}
}

type Child interface {
	SetParent(node GameNode)
	GetParent() GameNode
}

type GameComponent struct {
	Parent GameNode
}

func (this *GameComponent) SetParent(node GameNode) {
	this.Parent = node
}

func (this *GameComponent) GetParent() GameNode {
	return this.Parent
}

type GameNode interface {
	Component
	AddNode(node GameNode)
	AddComponent(component Component)
}

type Component interface {
	Child
	Initialize()
	Update(delta float32)
}

type Node struct {
	Transform
	GameComponent
	children   []GameNode
	components []Component
	Name       string
}

func (this *Node) Initialize() {
	for _, child := range this.children {
		child.Initialize()
	}
}

// func (this *Node) SetParent(node GameNode) {
// 	this.Parent = node
// }
//
// func (this *Node) GetParent() GameNode {
// 	return this.Parent
// }

func (this *Node) Update(delta float32) {
	if n, ok := this.Parent.(*Node); ok {
		this.model = n.GetUpdatedModel()
	}

	for _, component := range this.components {
		component.Update(delta)
	}

	for _, child := range this.children {
		child.Update(delta)
	}
}

// TODO make sure each node only has one of each type
func (this *Node) AddNode(node GameNode) {

	if n, ok := node.(Child); ok {

		n.SetParent(this)
	} else {
		fmt.Printf("No parent to set for child node: %s.\n", this.Name)
	}

	this.children = append(this.children, node)

	node.Initialize()
}

func (this *Node) AddComponent(component Component) {
	if n, ok := component.(Child); ok {

		n.SetParent(this)
	} else {
		fmt.Printf("No parent to set for child component.\n")
	}

	this.components = append(this.components, component)

	component.Initialize()
}
