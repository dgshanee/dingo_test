package components

import "encoding/json"

type TextComponent struct {
	Data      string `json:"data"`
	Component string `json:"component"`
}

func (txt TextComponent) Render() string {
	return ""
}

func TextFactory(jsonData json.RawMessage) (Component, error) {
	var txt TextComponent
	err := json.Unmarshal(jsonData, &txt);
	return txt, err;
}