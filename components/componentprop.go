package components

import "encoding/json"

type Component interface {
	Render() string
	GetHeight() int
	GetLine() int
	AddChild(component Component)
}

type ComponentProp struct {
	Component string      `json:"component"`
	Line      int         `json:"line"`
	Length    int         `json:"length"`
	Height    int         `json:"height"`
	ID        string      `json:"id"`
	Children  []Component `json:"children"`
}

func (cmp ComponentProp) AddChild(component Component) {
	cmp.Children = append(cmp.Children, component)
}

func (cmp ComponentProp) GetHeight() int {
	return cmp.Height
}

func (cmp ComponentProp) GetLine() int {
	return cmp.Line
}

type ComponentFactory func(jsonData json.RawMessage) (Component, error)

var componentFactories = map[string]ComponentFactory{
	"text": TextFactory,
}