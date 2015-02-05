package hallo

import (
	"net/http"
    "errors"
)

type fn func (*http.Request) (*Params,error)

type FuncManager struct {
    funcMap map[string] fn
}

func NewFuncManager() * FuncManager{
    return &FuncManager{ make(map[string]fn)}
}
func (f *FuncManager) addFunction(name string, newFun fn){
    f.funcMap[name]=newFun
}

func (f *FuncManager) getFunction(name string) (fn , error){

    if val, ok := f.funcMap[name]; ok {
        return val,nil
    }     
    return nil, errors.New(name+" Handling function Not Found")
}

func (f *FuncManager) remFunction(name string){
    delete(f.funcMap, name)
}

//----------------//
