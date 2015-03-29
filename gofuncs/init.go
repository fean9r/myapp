package hallo

import (
	"html/template"
	"net/http"
	"regexp"
    "github.com/fean9r/session"
    _ "github.com/fean9r/session/providers/memory"
     "errors"
)

var viewFuncManager = NewFuncManager()

var insertFuncManager = NewFuncManager()

var insertRequestManager = NewFuncManager()

var globalSessions *session.Manager

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
    
    viewFuncManager.addFunction("main", nil)
    viewFuncManager.addFunction("week", nil)
    viewFuncManager.addFunction("stats", nil)
    viewFuncManager.addFunction("day", dayPageFunc)
    viewFuncManager.addFunction("insertPage", insertPageFunc)
    viewFuncManager.addFunction("loginPage", loginPageFunc)
    viewFuncManager.addFunction("registerPage", registerPageFunc)

    insertFuncManager.addFunction("insertedValue", retriveInsertedValueData)
    insertFuncManager.addFunction("login", retriveInsertedLoginData)
    insertFuncManager.addFunction("register", retriveInsertedRegisterData)
    
    insertRequestManager.addFunction("insertedValue", processInsertedValueData)
    insertRequestManager.addFunction("login", processLoginData)
    insertRequestManager.addFunction("register", processRegisterData)

    err := errors.New("")
    // setting a new sessionManager
    globalSessions, err = session.NewManager("memory","gosessionid",3600) 
    if err != nil {
       reportValue("init",err)     
    }else {
        go globalSessions.GC()
    }
    
}


