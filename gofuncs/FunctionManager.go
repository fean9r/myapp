package hallo

import (
	"net/http"
    "errors"
)


type Params map[string] interface {}

type function func (http.ResponseWriter, *http.Request, *Params) error

type FuncManager struct {
    funcMap map[string] function
}

func NewFuncManager() * FuncManager{
    return &FuncManager{ make(map[string]function)}
}

func (f *FuncManager) addFunction(name string, newFun function){
    f.funcMap[name]=newFun
}

func (f *FuncManager) getFunction(name string) (function , error){
    if val, ok := f.funcMap[name]; ok {
        return val,nil
    }     
    return nil, errors.New(name+" Handling function Not Found")
}

func (f *FuncManager) remFunction(name string){
    delete(f.funcMap, name)
}

//----------------//
