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
	if err != nil{
		fmt.Println("Error loading components", err);
		return;
	}
	if parent != ""{
		if ok := DOMSlice.AddChildById(parent, component); !ok{
			fmt.Println("Error adding child");
			return;
		}

		fmt.Println(DOMSlice.GetSlice());
		
	} else {
		DOMSlice.AddComponent(component);
		fmt.Println(DOMSlice.GetSlice());
	}
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