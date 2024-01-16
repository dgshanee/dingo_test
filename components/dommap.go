package components

import "fmt"

type DomMap struct {
	domSlice []Component
	domMap   map[string]*Component
}

func (dm *DomMap) AddComponent(cmp Component) {
	dm.domSlice = append(dm.domSlice, cmp)
}

func (dm *DomMap) GetSlice() []Component {
	return dm.domSlice
}

func (dm *DomMap) GetComponentById(id string) (Component, bool) {
	res, ok := dm.domMap[id]
	fmt.Println("Component by id", *res)
	return *res, ok
}

func (dm *DomMap) AddChildById(id string, cmp Component) bool {
	parent, ok := dm.GetComponentById(id)
	if !ok {
		return false
	}

	parent.AddChild(cmp)
	return true
}
