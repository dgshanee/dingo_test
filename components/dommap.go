package components

import "errors"

type DomMap struct{
	domSlice []Component
	domMap map[string]Component
}

func (dm DomMap) getVal(id string) (Component,error){
	if res, ok := dm.domMap[id]; ok{
		return res, nil;
	}
	return nil, errors.New("Value not found in DOMMap");
}

func (dm DomMap) AddComponent(cmp Component){
	dm.domSlice = append(dm.domSlice, cmp);
}

func (dm DomMap) GetSlice() []Component{
	return dm.domSlice;
}

func (dm DomMap) GetComponentById(id string) (Component, bool){
	res, ok := dm.domMap[id];
	return res,ok;
}