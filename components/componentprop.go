package components

import "encoding/json"

type Component interface {
	Render() string
	GetHeight() int
	GetLine() int
	AddChild(component Component)
	GetId() (string, bool)
	SetID(string)
}

type ComponentProp struct {
	Component string      `json:"component"`
	Line      int         `json:"line"`
	Length    int         `json:"length"`
	Height    int         `json:"height"`
	ID        string      `json:"id"`
	Children  []Component `json:"children"`
}

func (cmp ComponentProp) GetId() (string, bool) {
	if cmp.ID != "" {
		return cmp.ID, true
	}
	return "", false
}

func (cmp *ComponentProp) AddChild(component Component) {
	cmp.Children = append(cmp.Children, component)
}

func (cmp ComponentProp) GetHeight() int {
	return cmp.Height
}

func (cmp ComponentProp) GetLine() int {
	return cmp.Line
}

func (cmp *ComponentProp) SetID(newID string) {
	cmp.ID = newID
}

type ComponentFactory func(jsonData json.RawMessage) (Component, error)

var componentFactories = map[string]ComponentFactory{
	"text": TextFactory,
	"body": BodyFactory,
}
