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

func LoadComponents() ([]components.Component, error){
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

func SaveComponent(component components.Component){
	DOMSlice, err := LoadComponents();
	fmt.Println("DOM",DOMSlice);
	if err != nil{
		fmt.Println("Error loading components", err);
		return;
	}
	DOMSlice = append(DOMSlice, component);
	marshalledComponent, err := json.Marshal(DOMSlice);
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

	for _,component := range DOMSlice{
		ln := component.GetLine();
		fmt.Println(ln);
	}
}

func Display(){
	
}