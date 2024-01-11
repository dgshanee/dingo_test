package components

import "encoding/json"

type TextComponent struct {
	ComponentProp
	Data string `json:"data"`
}

func (txt TextComponent) Render() string {
	return ""
}

func TextFactory(jsonData json.RawMessage) (Component, error) {
	var txt TextComponent
	err := json.Unmarshal(jsonData, &txt);
	return txt, err;
}