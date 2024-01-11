package dom

import (
	"dingo/components"
	"encoding/json"
	"fmt"
	"os"
)

var(
	lineMap map[int][]components.Component
	NUM_LINES int = 50
)

func LoadComponents() (*components.DomMap, error){
	data, err := os.ReadFile("structure.dingo");
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

func SaveComponent(component components.Component, parent string){
	DOMSlice, err := LoadComponents();
	fmt.Println("DOM",DOMSlice);
	if err != nil{
		fmt.Println("Error loading components", err);
		return;
	}
	if parent != ""{
		parentComponent, ok := DOMSlice.GetComponentById(parent);
		if !ok {
			fmt.Println("Invalid parent input");
			return;
		}

		parentComponent.AddChild(component);
	} else {
		DOMSlice.AddComponent(component);
	}
	fmt.Println("Slice", DOMSlice.GetSlice());
	marshalledComponent, err := json.Marshal(DOMSlice.GetSlice());
	if err != nil{
		fmt.Println("Error marshalling DOM", err);
		return;
	}

	err = os.WriteFile("structure.dingo", marshalledComponent, 0644);
	if err != nil{
		fmt.Println("Error writing to JSON", err);
		return;
	}
}

func populateLineMap(){
	lineMap = make(map[int][]components.Component);

	DOMSlice, err := LoadComponents();
	if err != nil{
		fmt.Println("Error reading dingo", err);
		return;
	}

	for _,component := range DOMSlice.GetSlice(){
		ln := component.GetLine();
		fmt.Println(ln);
	}
}

func Display(){
	
}