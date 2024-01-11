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