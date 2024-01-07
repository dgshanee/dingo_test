package components

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type Component interface {
	Render() string;
}

type ComponentFactory func(jsonData json.RawMessage) (Component, error); 

var componentFactories = map[string]ComponentFactory{
	"text":TextFactory,
}

func PopulateStruct(target interface{}, data map[string]string){
	v := reflect.ValueOf(target).Elem();

	for i:=0; i<v.NumField(); i++{
		field := v.Field(i);
		fieldName := v.Type().Field(i).Name;

		key := strings.ToLower(fieldName);

		if value, ok := data[key]; ok && field.CanSet(){
			field.SetString(value);
		}
	}
}

func UnmarshalJSONToComponent(jsonData []byte) ([]Component, error) {
	var rawMessages []json.RawMessage;
	if err := json.Unmarshal(jsonData, &rawMessages); err != nil{
		return nil, err;
	}
	type base struct {
			Type string `json:"component"`
	}

	var renderers []Component;


	for _,rawMessage := range rawMessages{
		var bs base;
		if err := json.Unmarshal([]byte(rawMessage), &bs); err != nil{
			return nil, err;
		}
		factory, ok := componentFactories[bs.Type];
		if !ok{
			fmt.Println("Unknown data type", bs.Type);
			continue;
		}

		component, err := factory(rawMessage);
		if err != nil{
			fmt.Println("Error getting component", err);
			continue;
		}

		renderers = append(renderers, component);
		
	}
	return renderers, nil;
}

func UnmarshalJSON[T any](jsonData []byte)(T, error){
	var t T;
	if err := json.Unmarshal(jsonData, &t); err != nil{
		return *new(T), err;
	}
	return t, nil;
}