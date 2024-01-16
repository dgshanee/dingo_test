package dom

import (
	"dingo/components"
	"encoding/json"
	"fmt"
	"os"

	"github.com/thanhpk/randstr"
)

var (
	lineMap   map[int][]components.Component
	NUM_LINES int = 50
	Body_id   string
)

func LoadComponents() (*components.DomMap, error) {
	data, err := os.ReadFile("structure.dingo")
	if err != nil {
		fmt.Println("Error reading JSON", err)
		return nil, err
	}

	DOMSlice, err := components.UnmarshalJSONToComponent(data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON", err)
		return nil, err
	}

	searchForBody(DOMSlice.GetSlice())

	return DOMSlice, nil
}

func searchForBody(comps []components.Component) {
	for _, comp := range comps {
		if comp.GetComponent() == "body" {
			data, ok := comp.GetId()
			if ok {
				Body_id = data
				return
			}
		}
	}
}

func SaveComponent(component components.Component, parent string) {
	DOMSlice, err := LoadComponents()
	fmt.Println("bodyid", Body_id)

	if err != nil {
		fmt.Println("Error loading components", err)
		return
	}
	if parent != "" {
		if ok := DOMSlice.AddChildById(parent, component); !ok {
			fmt.Println("Error adding child")
			return
		}

	} else {
		DOMSlice.AddComponent(component)
	}
	marshalledComponent, err := json.Marshal(DOMSlice.GetSlice())
	if err != nil {
		fmt.Println("Error marshalling DOM", err)
		return
	}

	err = os.WriteFile("structure.dingo", marshalledComponent, 0644)
	if err != nil {
		fmt.Println("Error writing to JSON", err)
		return
	}
}

func initBody() {
	var body = components.BodyComponent{
		components.ComponentProp{
			Component: "body",
		},
	}

	token := randstr.String(16)
	body.SetID(token)
	Body_id = token
	SaveComponent(&body, "")
}

func populateLineMap() {
	lineMap = make(map[int][]components.Component)

	DOMSlice, err := LoadComponents()
	if err != nil {
		fmt.Println("Error reading dingo", err)
		return
	}

	for _, component := range DOMSlice.GetSlice() {
		ln := component.GetLine()
		fmt.Println(ln)
	}
}

func Display() {
	fmt.Println("Displaying")
	LoadComponents()
}

func Clear() {
	clear := []string{}
	marshalledClear, err := json.Marshal(clear)
	if err != nil {
		fmt.Println("Error clearing DOM")
		return
	}

	err = os.WriteFile("structure.dingo", marshalledClear, 0644)
	if err != nil {
		fmt.Println("Error writing to file")
		return
	}

}
