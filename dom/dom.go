package dom

import (
	"dingo/components"
	"encoding/json"
	"fmt"
	"os"
)

func LoadComponents() ([]components.Component, error){
	data, err := os.ReadFile("structure.json");
	if err != nil{
		fmt.Println("Error reading JSON", err);
		return nil, err;
	}

	DOMSlice, err := components.UnmarshalJSONToComponent(data);
	if err != nil{
		fmt.Println("Error unmarshalling JSON", err);
		return nil, err;
	}

	return DOMSlice, nil;
}

func SaveComponent(component components.Component){
	DOMSlice, err := LoadComponents();
	if err != nil{
		fmt.Println("Error loading components", err);
		return;
	}
	fmt.Println(DOMSlice);
	DOMSlice = append(DOMSlice, component);
	fmt.Println(DOMSlice);
	marshalledComponent, err := json.Marshal(DOMSlice);
	if err != nil{
		fmt.Println("Error marshalling DOM", err);
		return;
	}

	err = os.WriteFile("structure.json", marshalledComponent, 0644);
	if err != nil{
		fmt.Println("Error writing to JSON", err);
		return;
	}
}