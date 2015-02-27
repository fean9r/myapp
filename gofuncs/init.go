package hallo

import (
	"html/template"
	"net/http"
	"regexp"
	"specificPages/"
)

var viewFuncManager = NewFuncManager()

var insertFuncManager = NewFuncManager()

var insertRequestManager = NewFuncManager()

var globalSessions *Manager

func first(args ...interface{}) interface{} {
    return args[0]
}

var templates = template.Must(template.ParseGlob("gtpl/*"))




func renderTemplate (w http.ResponseWriter, tmpl string, p  interface {} ) {
    err := templates.ExecuteTemplate(w, tmpl, p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}


func init() {
	
	addHandler("/insert/", makeHandler(insert))
    addHandler("/view/", makeHandler(view))
    validPath = regexp.MustCompile("^/("+ HandleFunctions +")/([a-zA-Z0-9]+)$")
    
    viewFuncManager.addFunction("main",nil)
    viewFuncManager.addFunction("week",nil)
    viewFuncManager.addFunction("stats",nil)
    viewFuncManager.addFunction("day",dayFunc)
    viewFuncManager.addFunction("insertPage",insertPageFunc)
    viewFuncManager.addFunction("loginPage",loginFunc)


    insertFuncManager.addFunction("insertedValue",retriveInsertedValueData)
    insertFuncManager.addFunction("login",retriveLoginData)
    
    
    insertRequestManager.addFunction("insertedValue",processInsertedValueData)
    insertRequestManager.addFunction("login",processLoginFunc)



    // setting a new sessionManager
    globalSessions = first( NewManager("memory","gosessionid",3600) ).(*Manager)


}


