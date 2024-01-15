package components

import "encoding/json"

type BodyComponent struct {
	ComponentProp
}

func (bd BodyComponent) Render() string {
	return ""
}

func BodyFactory(jsonData json.RawMessage) (Component, error) {
	var bd BodyComponent
	err := json.Unmarshal(jsonData, &bd)
	return &bd, err
}
